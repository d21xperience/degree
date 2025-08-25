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
        const ether = (wei / 10n ** 18n).toString(); // bagian integer ETH
        let remainder = wei % 10n ** 18n; // sisa dalam wei

        // Tambahkan 4 digit desimal dari remainder
        // Kita ambil 4 digit pertama dari sisa (karena 1e18 punya 18 digit, kita ambil 4)
        const decimalPart = ((remainder * 10000n) / 10n ** 18n).toString().padStart(4, '0');

        // Gabungkan integer dan desimal
        return `${ether}.${decimalPart}`;
    };

    const formatTimestamp = (timestamp) => {
        return new Date(timestamp * 1000).toLocaleString();
    };

    /**
     * Mengambil tahun dari input yang bisa berupa string tanggal, angka,
     * atau objek yang dapat dikonversi ke Date, lalu mengembalikan
     * tahun tersebut sebagai unsigned 32-bit integer (uint32).
     *
     * Mendukung format seperti:
     * - "Mon Jan 01 2024 07:00:00 GMT+0700 (Indochina Time)"
     * - "2023/2024"
     * - "2024"
     * - 2023
     * - Objek Date
     *
     * @param {string|number|Date} value - Nilai yang mengandung informasi tahun
     * @returns {number} Tahun sebagai unsigned 32-bit integer (uint32)
     * @throws {Error} Jika nilai tidak dapat diurai menjadi tanggal atau tahun yang valid
     *
     * @example
     * extractYearAsUint32("Mon Jan 01 2024 07:00:00 GMT+0700"); // 2024
     * extractYearAsUint32("2023/2024"); // 2023
     * extractYearAsUint32(2025); // 2025
     * extractYearAsUint32(new Date(2026, 0, 1)); // 2026
     */
    const extractYearAsUint32 = (value) => {
        if (value === null || value === undefined) {
            throw new Error('Nilai tidak boleh null atau undefined');
        }

        // Jika value sudah berupa angka
        if (typeof value === 'number') {
            if (!Number.isInteger(value) || value < 0 || value > 4294967295) {
                throw new Error('Angka harus dalam rentang uint32 (0 - 4294967295)');
            }
            return value >>> 0;
        }

        // Jika value adalah objek Date
        if (value instanceof Date) {
            const year = value.getFullYear();
            return year >>> 0;
        }

        // Konversi ke string
        const str = value.toString().trim();

        // Coba parse sebagai Date string (misal: "Mon Jan 01 2024...")
        const date = new Date(str);
        if (!isNaN(date.getTime())) {
            return date.getFullYear() >>> 0;
        }

        // Jika tidak bisa di-parse sebagai Date, coba ambil 4 digit pertama (misal: "2023/2024")
        const yearMatch = str.match(/^\d{4}/);
        if (yearMatch) {
            const year = parseInt(yearMatch[0], 10);
            return year >>> 0;
        }

        throw new Error('Tidak dapat mengekstrak tahun dari nilai: ' + str);
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
        formatTimestamp,
        extractYearAsUint32
    };
}
