import type { ShortenRequest, ShortenResponse, ApiError } from './types';

// Load from .env (Note the VITE_ prefix is required)
const API_URL = import.meta.env.VITE_GO_SHORTEN_URL || "http://localhost:3000/api/shorten";

export const shortenUrl = async (data: ShortenRequest): Promise<ShortenResponse> => {
    try {
        const response = await fetch(API_URL, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        });

        const responseData = await response.json();

        if (!response.ok) {
            const errorMsg = (responseData as ApiError).error || "Failed to shorten URL";
            throw new Error(errorMsg);
        }

        return responseData as ShortenResponse;
    } catch (error) {
        if (error instanceof Error) {
            throw error;
        }
        throw new Error("Network error. Is the Go backend running?");
    }
};