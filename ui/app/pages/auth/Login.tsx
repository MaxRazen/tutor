import React from 'react';
import LogoLink from '../../components/LogoLink';
import CardCentered from '../../components/CardCentered';

export default function Login () {
    return (
        <section>
            <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
                <div className="pb-6">
                    <LogoLink/>
                </div>

                <CardCentered>
                    <div>
                        <h1 className="text-xl font-bold leading-tight tracking-tight text-white md:text-2xl mb-6">
                            Welcome back
                        </h1>
                        <a
                            href="/auth/redirect/google"
                            className="btn-link"
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
                </CardCentered>
            </div>
        </section>
    )
}
