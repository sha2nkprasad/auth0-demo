import React from 'react'
import ReactDOM from 'react-dom/client'
import { Auth0Provider } from "@auth0/auth0-react";
import App from './App'
import './index.css'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <Auth0Provider
        domain={import.meta.env.VITE_AUTH0_DOMAIN}
        clientId={import.meta.env.VITE_AUTH0_CLIENT_ID}
        authorizationParams={{
            audience: import.meta.env.VITE_AUTH0_AUDIENCE,
            redirect_uri: 'http://localhost:5173/',
        }}
    >
        <React.StrictMode>
            <App />
        </React.StrictMode>
    </Auth0Provider>
)
