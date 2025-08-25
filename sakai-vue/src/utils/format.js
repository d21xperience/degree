/**
 * Format utilities for IPFS Admin
 */

/**
 * Format bytes to human readable string
 * @param {number} bytes
 * @param {number} decimals
 * @returns {string}
 */
export function formatBytes(bytes, decimals = 2) {
    if (bytes === 0) return '0 Bytes';

    const k = 1024;
    const dm = decimals < 0 ? 0 : decimals;
    const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];

    const i = Math.floor(Math.log(bytes) / Math.log(k));

    return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + ' ' + sizes[i];
}

/**
 * Format date to readable string
 * @param {Date|string|number} date
 * @returns {string}
 */
export function formatDate(date) {
    if (!date) return 'Unknown';

    const d = new Date(date);
    if (isNaN(d.getTime())) return 'Invalid date';

    return d.toLocaleString('en-US', {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit',
        hour12: true
    });
}

/**
 * Format peer ID to shorter version
 * @param {string} peerId
 * @returns {string}
 */
export function formatPeerId(peerId) {
    if (!peerId || peerId.length < 10) return peerId;
    return `${peerId.substring(0, 6)}...${peerId.substring(peerId.length - 4)}`;
}

/**
 * Format CID to shorter version
 * @param {string} cid
 * @returns {string}
 */
export function formatCid(cid) {
    if (!cid || cid.length < 10) return cid;
    return `${cid.substring(0, 6)}...${cid.substring(cid.length - 4)}`;
}

/**
 * Format duration in milliseconds to human readable string
 * @param {number} ms
 * @returns {string}
 */
export function formatDuration(ms) {
    if (ms < 1000) return `${ms}ms`;

    const seconds = Math.floor(ms / 1000);
    if (seconds < 60) return `${seconds}s`;

    const minutes = Math.floor(seconds / 60);
    const remainingSeconds = seconds % 60;
    if (minutes < 60) return `${minutes}m ${remainingSeconds}s`;

    const hours = Math.floor(minutes / 60);
    const remainingMinutes = minutes % 60;
    return `${hours}h ${remainingMinutes}m`;
}

/**
 * Format IPFS gateway URL
 * @param {string} cid
 * @param {string} path
 * @param {string} gateway
 * @returns {string}
 */
export function formatIpfsUrl(cid, path = '', gateway = 'https://ipfs.io') {
    if (!cid) return '';

    let url = `${gateway}/ipfs/${cid}`;
    if (path) {
        if (!path.startsWith('/')) path = `/${path}`;
        url += path;
    }

    return url;
}

/**
 * Format error message from various error types
 * @param {Error|string|any} error
 * @returns {string}
 */
export function formatError(error) {
    if (!error) return 'Unknown error';

    if (typeof error === 'string') return error;

    if (error.message) {
        // Clean up common error messages
        let msg = error.message;
        msg = msg.replace('Error: ', '');
        msg = msg.replace('error: ', '');
        return msg.charAt(0).toUpperCase() + msg.slice(1);
    }

    return JSON.stringify(error);
}

/**
 * Format boolean to Yes/No
 * @param {boolean} value
 * @returns {string}
 */
export function formatYesNo(value) {
    return value ? 'Yes' : 'No';
}

/**
 * Format number with thousands separator
 * @param {number} num
 * @returns {string}
 */
export function formatNumber(num) {
    if (typeof num !== 'number') return num;
    return num.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ',');
}

/**
 * Format IPFS API endpoint for display
 * @param {string} endpoint
 * @returns {string}
 */
export function formatApiEndpoint(endpoint) {
    if (!endpoint) return '';

    // Hide sensitive parts of the endpoint if needed
    return endpoint.replace('http://', '').replace('https://', '').replace('/api/v0', '');
}

export default {
    formatBytes,
    formatDate,
    formatPeerId,
    formatCid,
    formatDuration,
    formatIpfsUrl,
    formatError,
    formatYesNo,
    formatNumber,
    formatApiEndpoint
};
