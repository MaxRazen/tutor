import React from 'react';
import { createRoot } from 'react-dom/client';
import {
    createBrowserRouter,
    RouterProvider,
  } from 'react-router-dom';
import Home from './pages/Home';
import About from './pages/About';
import Login from './pages/Login';
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
        element: <Login/>,
    },
]);

const root = createRoot(document.getElementById("app")!);

root.render(
    <RouterProvider router={router} />
);