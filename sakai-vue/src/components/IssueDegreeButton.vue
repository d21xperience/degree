<template>
    <Button icon="pi pi-file-export" @click="handleSubmit" :loading="isLoading" v-tooltip.bottom="'Verifikasi ijazah'" />
</template>

<script setup>
import { useSCService } from '@/composables/useSCService';
import DegreeContractABI from '@/VerifikasiIjazahABI.json';
import { ethers, keccak256, toUtf8Bytes } from 'ethers';
import { onMounted, ref } from 'vue';
const { createMetamaskConnected, createSCIjazah } = useSCService();
const contractAddress = '0xdc64a140aa3e981100a9beca4e685f962f0cf6c9';
const props = defineProps({
    degreeData: Object, // { name, nisn, graduationYear, major }
    sekolah: String,
    ipfsUrl: String,
    transcript: Object, // { subjects: ["B. Indonesia", "Matematika"], grades: [85, 90] }
    tahunAjaranId: String
});
const emit = defineEmits(['submit']);
// State
const isLoading = ref(false);
const contract = ref(null);

// Inisialisasi kontrak
const loadContract = async () => {
    if (window.ethereum) {
        const provider = new ethers.BrowserProvider(window.ethereum);
        const signer = await provider.getSigner();
        contract.value = new ethers.Contract(contractAddress, DegreeContractABI, signer);
    } else {
        alert('Metamask tidak ditemukan.');
    }
};

// Fungsi untuk membuat hash ijazah (gunakan string JSON untuk konsistensi)
const generateDegreeHash = (data) => {
    const stringified = JSON.stringify(data);
    return keccak256(toUtf8Bytes(stringified)); // ethers v6
};

const handleSubmit = async () => {
    // let txtHas = '1232443';
    // let degreeHash = '123244dfdfdf';
    // saveToBackend(txtHas, degreeHash);

    isLoading.value = true;
    try {
        if (!contract.value) await loadContract();

        const degreeHash = generateDegreeHash(props.degreeData);
        const issueDate = Math.floor(Date.now() / 1000);

        // Validasi & parsing subjects dan grades
        let subjects = props.transcript?.subjects;
        let grades = props.transcript?.grades;

        if (!Array.isArray(subjects) || !Array.isArray(grades)) {
            throw new Error('Data transcript tidak valid. Pastikan subjects dan grades berupa array.');
        }

        grades = grades.map((n) => parseInt(n));

        if (subjects.length !== grades.length) {
            throw new Error('Jumlah mata pelajaran dan nilai tidak cocok.');
        }

        const gasEstimate = await contract.value.issueDegree.estimateGas(degreeHash, props.sekolah, issueDate, props.ipfsUrl, subjects, grades);

        const proceed = confirm(`Biaya gas kira-kira: ${ethers.formatUnits(gasEstimate, 'gwei')} Gwei. Lanjutkan?`);
        if (!proceed) return;

        const tx = await contract.value.issueDegree(degreeHash, props.sekolah, issueDate, props.ipfsUrl, subjects, grades);

        await tx.wait();
        alert('Ijazah berhasil diverifikasi di blockchain!');
        emit('submit');
        saveToBackend(tx.hash, degreeHash);
    } catch (err) {
        console.error(err);
        alert(`Gagal memproses: ${err.message}`);
    } finally {
        isLoading.value = false;
    }
};
// const handleBatchSubmit = async () => {
//     isLoading.value = true;
//     try {
//         if (!contract.value) await loadContract();

//         for (const student of studentsData) {
//             // Buat degreeHash dari data (hash dari ipfs + sekolah + waktu sekarang)
//             const degreeHash = generateDegreeHash(student); // pastikan fungsi ini konsisten

//             const issueDate = Math.floor(Date.now() / 1000);

//             // Validasi input
//             if (!Array.isArray(student.subjects) || !Array.isArray(student.grades) || student.subjects.length !== student.grades.length) {
//                 console.warn('Data tidak valid:', student);
//                 continue; // Skip data tidak valid
//             }

//             console.log('Mengirim untuk:', student.sekolah, degreeHash);

//             const tx = await contract.value.issueDegree(degreeHash, student.sekolah, issueDate, student.ipfsUrl, student.subjects, student.grades);

//             await tx.wait(); // Tunggu satu per satu
//             console.log('Sukses simpan:', degreeHash);
//         }

//         alert('Semua data siswa berhasil dikirim ke smart contract.');
//     } catch (err) {
//         console.error('Gagal:', err);
//         alert(`Terjadi error: ${err.message}`);
//     } finally {
//         isLoading.value = false;
//     }
// };

// const contract = ref(null);

onMounted(async () => {
    try {
        if (window.ethereum) {
            await window.ethereum.request({ method: 'eth_requestAccounts' });
            const provider = new ethers.BrowserProvider(window.ethereum);
            const signer = await provider.getSigner();

            // const contractAddress = '0x700b6A60ce7EaaEA56F065753d8dcB9653dbAD35'; // HARUS STRING BUKAN undefined/objek
            contract.value = new ethers.Contract(contractAddress, DegreeContractABI, signer);
            if (contract.value) {
                createMetamaskConnected(true);
            }
        } else {
            alert('Metamask tidak ditemukan. Harap instal terlebih dahulu.');
        }
    } catch (err) {
        console.error('Gagal menginisialisasi kontrak:', err);
    }
});
// 🗄️ Simpan ke backend
const saveToBackend = async (txHash, degreeHash) => {
    if (!props.degreeData) {
        console.error('degreeData is undefined');
        return;
    }
    // console.log(props.degreeData);
    const penerimaIjazah = props.degreeData[0];
    let degreeData = {
        degreeHash: degreeHash,
        txHash: txHash,
        ipfsUrl: props.ipfsUrl,
        bcType: 'Ethereum',
        linkBcExplorer: 'http://localhost:26000',
        tahunAjaranId: props?.tahunAjaranId,
        sekolah_id: penerimaIjazah?.sekolahId,
        ijazah: {
            peserta_didik_id: penerimaIjazah?.pesertaDidikId,
            nama: penerimaIjazah?.nama,
            nis: penerimaIjazah?.nis,
            nisn: penerimaIjazah?.nisn,
            npsn: penerimaIjazah?.npsn,
            nomor_ijazah: penerimaIjazah?.nomorIjazah,
            tempat_lahir: penerimaIjazah?.tempatLahir,
            tanggal_lahir: penerimaIjazah?.tanggalLahir,
            nama_ortuwali: penerimaIjazah?.namaOrtuWali,
            paket_keahlian: penerimaIjazah?.paketKeahlian,
            kabupatenkota: penerimaIjazah?.kabupatenKota,
            provinsi: penerimaIjazah?.provinsi,
            program_keahlian: penerimaIjazah?.programKeahlian,
            sekolah_penyelenggara_ujian_us: penerimaIjazah?.sekolahPenyelenggaraUjianUs,
            sekolah_penyelenggara_ujian_un: penerimaIjazah?.sekolahPenyelenggaraUjianUn,
            asal_sekolah: penerimaIjazah?.asalSekolah,
            tempat_ijazah: penerimaIjazah?.tempatIjazah,
            tanggal_ijazah: penerimaIjazah?.tanggalIjazah
        },
        transkrip: { subjects: props.transcript?.subjects, grades: props.transcript?.grades }
    };

    createSCIjazah(degreeData);
};
</script>
