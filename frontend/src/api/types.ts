export interface ShortenRequest {
    url: string;
    short?: string; // Optional custom short
    expiry: number; // Sent as HOURS
}

export interface ShortenResponse {
    url: string;
    short: string;
    expiry: number;           // Received as HOURS
    rate_limit: number;       // Requests remaining
    rate_limit_reset: number; // Received as MINUTES
}

export interface ApiError {
    error: string;
}