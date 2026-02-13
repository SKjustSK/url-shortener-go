import React, { useState } from 'react';
import { shortenUrl } from '../api/api';
import type { ShortenResponse } from '../api/types';

const ShortenerForm: React.FC = () => {
    const [url, setUrl] = useState('');
    const [customShort, setCustomShort] = useState('');
    const [expiry, setExpiry] = useState<number>(24);
    
    const [result, setResult] = useState<ShortenResponse | null>(null);
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState(false);

    // Get Domain from Env & Clean it
    const rawDomain = import.meta.env.VITE_GO_DOMAIN || "localhost:3000";
    const displayDomain = rawDomain.replace(/^https?:\/\//, '');

    // Helper: Ensure the link is clickable
    const formatLink = (link: string) => {
        if (!link.startsWith('http')) return `http://${link}`;
        return link;
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setLoading(true);
        setError(null);
        setResult(null);

        if (!url) {
            setError("URL is required");
            setLoading(false);
            return;
        }

        try {
            const response = await shortenUrl({
                url: url,
                short: customShort,
                expiry: expiry 
            });
            
            setResult(response);
        } catch (err: any) {
            setError(err.message || "Something went wrong");
        } finally {
            setLoading(false);
        }
    };

    return (
        // Added 'relative' here to position the GitHub icon
        <div className="bg-white p-8 rounded-xl shadow-lg w-full max-w-md border border-gray-100 relative">
            
            {/* --- GitHub Logo Link --- */}
            <a 
                href="https://github.com/SKjustSK/url-shortener-go" 
                target="_blank" 
                rel="noopener noreferrer"
                className="absolute top-5 right-5 text-gray-400 hover:text-gray-900 transition-colors duration-200"
                title="View Source on GitHub"
            >
                <svg viewBox="0 0 16 16" fill="currentColor" className="w-6 h-6">
                    <path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z"/>
                </svg>
            </a>
            {/* ------------------------ */}

            <div className="text-center mb-8">
                <h2 className="text-3xl font-extrabold text-gray-900">URL Shortener</h2>
                <p className="text-gray-500 mt-2 text-sm">Enter a long link to make it short</p>
            </div>
            
            <form onSubmit={handleSubmit} className="space-y-5">
                <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1">Destination URL</label>
                    <input 
                        type="url" 
                        required 
                        placeholder="https://example.com/very-long-link"
                        value={url}
                        onChange={(e) => setUrl(e.target.value)}
                        className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none transition shadow-sm"
                    />
                </div>

                <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1">Custom Short (Optional)</label>
                    <div className="flex rounded-md shadow-sm">
                        <span className="inline-flex items-center px-3 rounded-l-md border border-r-0 border-gray-300 bg-gray-50 text-gray-500 text-sm select-none whitespace-nowrap">
                           {displayDomain}/
                        </span>
                        <input 
                            type="text" 
                            placeholder="my-link"
                            value={customShort}
                            onChange={(e) => setCustomShort(e.target.value)}
                            className="flex-1 min-w-0 block w-full px-4 py-3 border border-gray-300 rounded-r-md focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm outline-none"
                        />
                    </div>
                </div>

                <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1">Expiry (Hours)</label>
                    <input 
                        type="number" 
                        min="1"
                        value={expiry}
                        onChange={(e) => setExpiry(Number(e.target.value))}
                        className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none transition shadow-sm"
                    />
                </div>

                <button 
                    type="submit" 
                    disabled={loading}
                    className={`w-full py-3 px-4 rounded-lg text-white font-bold text-lg shadow-md transition duration-200 transform hover:-translate-y-0.5
                        ${loading ? 'bg-indigo-300 cursor-not-allowed' : 'bg-indigo-600 hover:bg-indigo-700'}`}
                >
                    {loading ? "Processing..." : "Shorten URL"}
                </button>
            </form>

            {error && (
                <div className="mt-6 p-4 bg-red-50 text-red-700 border-l-4 border-red-500 rounded-r-md text-sm">
                    <p className="font-bold">Error</p>
                    <p>{error}</p>
                </div>
            )}

            {result && (
                <div className="mt-8 bg-green-50 border border-green-200 rounded-xl p-5 shadow-inner">
                    <div className="flex items-center justify-between mb-4">
                        <h3 className="text-lg font-bold text-green-800">Success!</h3>
                        <span className="px-2 py-1 bg-green-200 text-green-800 text-xs font-bold rounded uppercase">Active</span>
                    </div>
                    
                    <div className="bg-white p-3 rounded border border-green-100 flex justify-between items-center mb-3">
                        <a href={formatLink(result.short)} target="_blank" rel="noreferrer" className="text-indigo-600 font-bold hover:underline text-lg truncate">
                            {result.short}
                        </a>
                        <button 
                            onClick={() => navigator.clipboard.writeText(formatLink(result.short))}
                            className="ml-2 text-xs bg-gray-100 hover:bg-gray-200 text-gray-600 px-2 py-1 rounded border border-gray-300"
                        >
                            Copy
                        </button>
                    </div>

                    <div className="border-t border-green-200 pt-3 flex justify-between text-xs text-green-800 font-medium">
                        <div className="flex flex-col">
                            <span className="text-green-600 uppercase tracking-wider text-[10px]">Rate Limit</span>
                            <span>{result.rate_limit} remaining</span>
                        </div>
                        <div className="flex flex-col text-right">
                            <span className="text-green-600 uppercase tracking-wider text-[10px]">Resets In</span>
                            <span>{result.rate_limit_reset} min</span>
                        </div>
                    </div>
                </div>
            )}
        </div>
    );
};

export default ShortenerForm;