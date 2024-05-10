import React from 'react';

export default function Login () {
    return (
        <section className="bg-gray-50 dark:bg-gray-900">
            <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
                <a
                    href="/"
                    className="flex items-center mb-6 text-2xl font-semibold text-gray-900 dark:text-white"
                    >
                    <img
                        className="w-8 h-8 mr-2"
                        src="/assets/logo-white.svg"
                        alt="logo"
                    />
                    AI Tutor <sup className="text-gray-300 pl-2 font-normal text-xs">by Hexnet</sup>
                </a>
                <div className="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
                <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
                    <h1 className="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
                        Welcome back
                    </h1>
                    <a
                        href="/auth/redirect/google"
                        className="w-full p-4 border border-gray-700 rounded-lg flex items-center justify-center hover:border-gray-600 hover:bg-gray-700"
                    >
                        <img
                            src="/assets/google-icon.svg"
                            width="20"
                            height="20"
                            alt="Google Auth"
                            className="mr-2"
                        ></img>
                        Log In with Google
                    </a>
                </div>
                </div>
            </div>
        </section>
    )
}