<template>
  <div class="p-4 space-y-4">
    <input type="file" @change="onFileChange" accept="application/pdf" />
    <input v-model="sekolah" placeholder="Nama Sekolah" class="border p-2" />
    <button @click="uploadAndHash" :disabled="!pdfFile">1. Upload & Hash</button>
    <button @click="sendToBlockchain" :disabled="!fileHash">2. Simpan ke Blockchain</button>
    <div v-if="txHash">✅ Transaksi berhasil: {{ txHash }}</div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import axios from 'axios'
import { ethers } from 'ethers'
import contractAbi from '@/VerifikasiIjazahABI.json'

const pdfFile = ref(null)
const fileHash = ref('')
const sekolah = ref('')
const txHash = ref('')

// Ganti dengan alamat kontrak Anda
const contractAddress = '0xYourContractAddress'

const onFileChange = (e) => {
  pdfFile.value = e.target.files[0]
}

const uploadAndHash = async () => {
  const formData = new FormData()
  formData.append('file', pdfFile.value)

  const res = await axios.post('http://localhost:8080/ijazah/hash', formData)
  fileHash.value = res.data.file_hash
  alert(`Hash: ${fileHash.value}`)
}

const sendToBlockchain = async () => {
  if (!window.ethereum) {
    alert('MetaMask tidak tersedia')
    return
  }

  const provider = new ethers.BrowserProvider(window.ethereum)
  const signer = await provider.getSigner()
  const contract = new ethers.Contract(contractAddress, contractAbi, signer)

  const timestamp = Math.floor(Date.now() / 1000)

  const tx = await contract.issueDegree(
    fileHash.value,
    sekolah.value,
    timestamp,
    '', // Kosongkan dulu jika belum pakai IPFS
    [], // Mata pelajaran
    []  // Nilai
  )

  const receipt = await tx.wait()
  txHash.value = receipt.transactionHash
}
</script>
