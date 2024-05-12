import React from 'react';
import {AuthCallbackPageData} from '../../types.d';
import {Link, useNavigate} from 'react-router-dom';
import LogoLink from '../../components/LogoLink';
import CardCentered from '../../components/CardCentered';
import Alert from '../../components/Alert';
import store from '../../store';
import app from '../../Application';

export default function Callback() {
    app.setTitle('Auth');
    const navigate = useNavigate();
    const data: AuthCallbackPageData = window.PageData;

    const success = data.authorized;

    if (success) {
        store.setUser(data.user);
        store.setAccessToken(data.accessToken);
        setTimeout(() => navigate('/'), 1000);
    }

    const backLink = success
        ? <Link to={'/'} className="btn-link">Go to Home page</Link>
        : <Link to={'/login'} className="btn-link">Go back to Login page</Link>

    return (
        <section>
            <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
                <div className="py-8">
                    <LogoLink/>
                </div>

                <CardCentered>
                    <div>
                        <h1 className="text-base font-medium leading-tight tracking-tight text-white md:text-lg mb-2">
                            {success ? 'You\'ve been authenticated successuflly. Redirecting...' : 'Authentification failed'}
                        </h1>

                        {data.alert ? (
                            <Alert type={data.alert.type} message={data.alert.message}/>
                        ) : ''}

                        <div className="mt-8">
                            {backLink}
                        </div>
                    </div>
                </CardCentered>
            </div>
        </section>
    )
}
