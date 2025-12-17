import {useState, useEffect, useRef} from 'react';
import {scanDomain} from './api';
import {AnalyzeResponse, Status} from './types';
import './App.css';

function App() {
	const [domain, setDomain] = useState('');
	const [loading, setLoading] = useState(false);
	const [result, setResult] = useState<AnalyzeResponse | null>(null);
	const [error, setError] = useState<string | null>(null);
	const [polling, setPolling] = useState(false);
	const pollingIntervalRef = useRef<number | null>(null);

	// Cleanup polling on unmount
	useEffect(() => {
		return () => {
			if (pollingIntervalRef.current) {
				clearInterval(pollingIntervalRef.current);
			}
		};
	}, []);

	const stopPolling = () => {
		if (pollingIntervalRef.current) {
			clearInterval(pollingIntervalRef.current);
			pollingIntervalRef.current = null;
		}
		setPolling(false);
	};

	const startPolling = (domainToPoll: string) => {
		stopPolling();
		setPolling(true);

		const poll = async () => {
			try {
				// Always use cached results for polling (startNew=false)
				const response = await scanDomain(domainToPoll, false);
				setResult(response);
				setError(null);

				// Stop polling if we have a final result
				if (response.status === Status.READY || response.status === Status.ERROR) {
					stopPolling();
					setLoading(false);
				}
			} catch (err) {
				setError(err instanceof Error ? err.message : 'An error occurred');
				stopPolling();
				setLoading(false);
			}
		};

		// Poll immediately, then every 5 seconds
		poll();
		pollingIntervalRef.current = window.setInterval(poll, 5000);
	};

	const handleScan = async (startNew: boolean = false) => {
		if (!domain.trim()) {
			setError('Please enter a domain');
			return;
		}

		setLoading(true);
		setError(null);
		setResult(null);

		try {
			const response = await scanDomain(domain.trim(), startNew);
			setResult(response);

			// If the scan is in progress, start polling
			if (response.status === Status.DNS || response.status === Status.IN_PROGRESS) {
				startPolling(domain.trim());
			} else {
				setLoading(false);
			}
		} catch (err) {
			setError(err instanceof Error ? err.message : 'Failed to scan domain');
			setLoading(false);
		}
	};

	const getGradeColor = (grade?: string) => {
		if (!grade) return '#6b7280';
		const gradeUpper = grade.toUpperCase();
		if (gradeUpper.includes('A+') || gradeUpper === 'A') return '#000000'; // black for best
		if (gradeUpper === 'B') return '#374151'; // dark grey
		if (gradeUpper === 'C') return '#6b7280'; // medium grey
		if (gradeUpper === 'D' || gradeUpper === 'E') return '#9ca3af'; // light grey
		if (gradeUpper === 'F') return '#d1d5db'; // very light grey
		return '#6b7280'; // gray
	};

	const getGradeTextColor = (grade?: string) => {
		if (!grade) return '#ffffff';
		const gradeUpper = grade.toUpperCase();
		// Dark backgrounds get white text, light backgrounds get black text
		if (gradeUpper.includes('A+') || gradeUpper === 'A' || gradeUpper === 'B') return '#ffffff';
		if (gradeUpper === 'C') return '#ffffff';
		return '#000000'; // Light backgrounds get black text
	};

	const getStatusBadge = (status: string) => {
		switch (status) {
			case Status.READY:
				return {text: 'Ready', color: '#000000'};
			case Status.IN_PROGRESS:
				return {text: 'In Progress', color: '#374151'};
			case Status.DNS:
				return {text: 'DNS Lookup', color: '#6b7280'};
			case Status.ERROR:
				return {text: 'Error', color: '#9ca3af'};
			default:
				return {text: status, color: '#6b7280'};
		}
	};

	return (
		<div className="app">
			<div className="container">
				<header className="header">
					<h1 className="title">
						<span className="title-icon">🔒</span>
						TLS Protection Checker
					</h1>
					<p className="subtitle">Truora - josegabjimenez</p>
				</header>

				<div className="search-section">
					<div className="input-group">
						<input
							type="text"
							className="domain-input"
							placeholder="Enter domain (e.g., google.com)"
							value={domain}
							onChange={(e) => setDomain(e.target.value)}
							onKeyDown={(e) => {
								if (e.key === 'Enter' && !loading) {
									handleScan();
								}
							}}
							disabled={loading}
						/>
						<div className="button-group">
							<button className="btn btn-primary" onClick={() => handleScan(false)} disabled={loading || !domain.trim()}>
								{loading ? (
									<>
										<span className="spinner"></span>
										Scanning...
									</>
								) : (
									'Scan Domain'
								)}
							</button>
							<button className="btn btn-secondary" onClick={() => handleScan(true)} disabled={loading || !domain.trim()}>
								New Scan
							</button>
						</div>
					</div>
				</div>

				{error && (
					<div className="error-message">
						<span className="error-icon">⚠️</span>
						{error}
					</div>
				)}

				{result && (
					<div className="results-section">
						<div className="result-card">
							<div className="result-header">
								<div className="domain-info">
									<h2 className="result-domain">{result.host}</h2>
									<span
										className="status-badge"
										style={{backgroundColor: getStatusBadge(result.status).color + '20', color: getStatusBadge(result.status).color}}
									>
										{getStatusBadge(result.status).text}
									</span>
								</div>
								{polling && (
									<div className="polling-indicator">
										<span className="pulse-dot"></span>
										Polling for updates...
									</div>
								)}
							</div>

							{result.statusMessage && <p className="status-message">{result.statusMessage}</p>}

							{result.endpoints && result.endpoints.length > 0 && (
								<div className="endpoints-section">
									<h3 className="endpoints-title">Endpoints</h3>
									<div className="endpoints-grid">
										{result.endpoints.map((endpoint, index) => (
											<div key={index} className="endpoint-card">
												<div className="endpoint-header">
													<div className="endpoint-ip">{endpoint.ipAddress}</div>
													{endpoint.grade && (
														<div
															className="grade-badge"
															style={{
																backgroundColor: getGradeColor(endpoint.grade),
																color: getGradeTextColor(endpoint.grade),
															}}
														>
															{endpoint.grade}
														</div>
													)}
												</div>
												{endpoint.serverName && (
													<div className="endpoint-detail">
														<span className="detail-label">Server:</span>
														<span className="detail-value">{endpoint.serverName}</span>
													</div>
												)}
												{endpoint.progress !== undefined && endpoint.progress < 100 && (
													<div className="progress-bar">
														<div className="progress-fill" style={{width: `${endpoint.progress === -1 ? 0 : endpoint.progress}%`}}></div>
														<span
															className="progress-text"
															style={{
																color: (endpoint.progress === -1 ? 0 : endpoint.progress) >= 50 ? '#ffffff' : '#374151',
															}}
														>
															{endpoint.progress === -1 ? 0 : endpoint.progress}%
														</span>
													</div>
												)}
												<div className="endpoint-features">
													{endpoint.isExceptional && <span className="feature-badge exceptional">⭐ Exceptional</span>}
													{endpoint.hasWarnings && <span className="feature-badge warning">⚠️ Warnings</span>}
												</div>
											</div>
										))}
									</div>
								</div>
							)}
						</div>
					</div>
				)}

				{!result && !loading && !error && (
					<div className="empty-state">
						<div className="empty-icon">🔍</div>
						<p className="empty-text">Enter a domain above to check its TLS protection</p>
					</div>
				)}
			</div>
		</div>
	);
}

export default App;
