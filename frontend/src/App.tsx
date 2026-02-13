import ShortenerForm from './components/ShortenerForm'

function App() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-indigo-100 to-purple-100 flex flex-col items-center justify-center p-4">
      
      {/* Main Content */}
      <div className="flex-grow flex items-center justify-center w-full">
        <ShortenerForm />
      </div>

      {/* Footer / Render Warning */}
      <footer className="mt-8 text-center max-w-md">
        <div className="flex items-center justify-center gap-2 text-indigo-600 mb-1">
          <svg 
            xmlns="http://www.w3.org/2000/svg" 
            className="h-5 w-5 animate-pulse" 
            viewBox="0 0 20 20" 
            fill="currentColor"
          >
            <path fillRule="evenodd" d="M11.3 1.046A1 1 0 0112 2v5h4a1 1 0 01.82 1.573l-7 10A1 1 0 018 18v-5H4a1 1 0 01-.82-1.573l7-10a1 1 0 011.12-.38z" clipRule="evenodd" />
          </svg>
          <span className="font-semibold text-sm uppercase tracking-wider">System Status</span>
        </div>
        
        <p className="text-gray-500 text-xs leading-relaxed">
          Note: This project is hosted on a free instance. If the service has been inactive, 
          it may take <span className="font-medium text-indigo-700">50-60 seconds</span> for 
          the first request to process while the server wakes up.
        </p>
      </footer>

    </div>
  )
}

export default App