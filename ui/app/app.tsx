import React from 'react';
import { createRoot } from 'react-dom/client';
import {
    createBrowserRouter,
    RouterProvider,
} from 'react-router-dom';
import Home from './pages/Home';
import About from './pages/About';
import AuthLogin from './pages/auth/Login';
import AuthCallback from './pages/auth/Callback';
import '../styles/main.scss';

const router = createBrowserRouter([
    {
        path: '/',
        element: <Home/>,
    },
    {
        path: 'about',
        element: <About/>,
    },
    {
        path: '/login',
        element: <AuthLogin/>,
    },
    {
        path: '/auth/callback/:provider',
        element: <AuthCallback/>,
    },
]);

const root = createRoot(document.getElementById('app')!);

root.render(
    <RouterProvider router={router} />
);
