// Type definitions for SSL Labs API responses

export type AnalyzeResponse = {
	host: string;
	port: number;
	protocol: string;
	isPublic: boolean;
	status: string;
	statusMessage?: string;
	startTime: number;
	testTime?: number;
	engineVersion: string;
	criteriaVersion: string;
	endpoints?: Endpoint[];
};

export type Endpoint = {
	ipAddress: string;
	serverName?: string;
	statusMessage?: string;
	grade?: string;
	gradeTrustIgnored?: string;
	hasWarnings: boolean;
	isExceptional: boolean;
	progress?: number;
	duration?: number;
	delegation: number;
};

export const Status = {
	DNS: 'DNS',
	ERROR: 'ERROR',
	IN_PROGRESS: 'IN_PROGRESS',
	READY: 'READY',
} as const;
