import type {AnalyzeResponse} from './types';

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

export async function scanDomain(domain: string, startNew: boolean = false): Promise<AnalyzeResponse> {
	const url = `${API_BASE_URL}/api/scan?domain=${encodeURIComponent(domain)}&new=${startNew}`;

	const response = await fetch(url);

	if (!response.ok) {
		const error = await response.json().catch(() => ({error: 'Failed to scan domain'}));
		throw new Error(error.error || 'Failed to scan domain');
	}

	return response.json();
}
