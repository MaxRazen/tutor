import React from 'react';
import { Link } from 'react-router-dom';

export default function Home () {
    return (
        <>
            <div className="container py-10 mx-auto">
                <h1>This is home page</h1>
                <Link to={'login'}>Login</Link>
                <Link to={'about'}>About</Link>
            </div>
        </>
    )
}