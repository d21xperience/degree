export function useUtils() {
    const formatterDateID = (tanggalRaw) => {
        if (!tanggalRaw) return '-';

        // Normalize format: Ganti spasi ke 'T' jika diperlukan
        let normalizedDate = tanggalRaw.replace(' ', 'T');

        // Buat objek Date
        const tanggal = new Date(normalizedDate);

        // Cek validitas
        if (isNaN(tanggal)) return '-';

        const formatter = new Intl.DateTimeFormat('id-ID', {
            day: '2-digit',
            month: 'long',
            year: 'numeric'
        });

        return formatter.format(tanggal);
    };
    const ringkasHash = (hash, awal = 6, akhir = 6) => {
        if (!hash || hash.length < awal + akhir + 2) return hash;
        return `${hash.slice(0, awal + 2)}...${hash.slice(-akhir)}`;
    };
    const getWebsiteUrl = (url) => {
        if (!url.startsWith('http://') && !url.startsWith('https://')) {
            return `https://${url}`; // Tambahkan https jika belum ada
        }
        return url;
    };

    const shortenAddress = (address) => {
        return `${address.substring(0, 6)}...${address.substring(address.length - 4)}`;
    };

    const shortenHash = (hash) => {
        return `${hash.substring(0, 8)}...`;
    };

const formatBalance = (balance) => {
    // Handle BigInt
    const wei = typeof balance === 'bigint' ? balance : BigInt(Math.floor(Number(balance) || 0));

    // Pisahkan bagian integer dan desimal dari wei
    const ether = (wei / 10n**18n).toString(); // bagian integer ETH
    let remainder = wei % 10n**18n; // sisa dalam wei

    // Tambahkan 4 digit desimal dari remainder
    // Kita ambil 4 digit pertama dari sisa (karena 1e18 punya 18 digit, kita ambil 4)
    const decimalPart = (remainder * 10000n / 10n**18n).toString().padStart(4, '0');

    // Gabungkan integer dan desimal
    return `${ether}.${decimalPart}`;
};

    const formatTimestamp = (timestamp) => {
        return new Date(timestamp * 1000).toLocaleString();
    };

    return {
        // tingkatPendidikanOptions,
        // jurusanOptions,
        formatterDateID,
        ringkasHash,
        getWebsiteUrl,
        formatBalance,
        shortenHash,
        shortenAddress,
        formatTimestamp
    };
}
