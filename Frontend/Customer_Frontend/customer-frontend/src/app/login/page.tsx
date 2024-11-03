'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { useNavBarStore, useUserStore } from '../../../store/store';
import { setCookie } from 'cookies-next';

const Login = () => {
    const [username, setUsername] = useState('');
    const [logInEmail, setLogInEmail] = useState('');
    const [signUpEmail, setSignUpEmail] = useState('');
    const [logInPassword, setLogInPassword] = useState('');
    const [signUpPassword, setSignUpPassword] = useState('');
    const [renderErrorUsername, setRenderErrorUsername] = useState(false);
    const [signUpServerError, setSignUpServerError] = useState('');
    const [logInServerError, setLogInServerError] = useState('');
    const [loadingButton, setLoadingButton] = useState('');
    const router = useRouter();
    const { setIsNavBarOpen } = useNavBarStore() as { setIsNavBarOpen: (value: boolean) => void };
    const { setEmail } = useUserStore() as { setUsername: (value: string) => void, setEmail: (value: string) => void };

    useEffect(() => {
        setIsNavBarOpen(false);

        // Cleanup
        return () => {
            setIsNavBarOpen(true);
        }
    }, [setIsNavBarOpen]);

    const handleLogin = async (e: React.FormEvent) => {
        e.preventDefault();

        
        // Verify all fields are filled
        if (logInEmail.length === 0 || logInPassword.length === 0) {
            return;
        }
        setLoadingButton('Login');
        
        const url = process.env.API_GATEWAY || 'https://xqtvgb1k-8080.asse.devtunnels.ms';
        console.log('API Gateway:', url);
        const uri = '/rest/v1/user/login';
        const endpoint = url + uri;
        try {
            // API call here - POST request to /rest/v1/user/login, with email and password in header
            const response = await fetch(endpoint, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    email: logInEmail,
                    password: logInPassword,
                }),
            });
            const data = await response.json();
            
            // If successful, store token in local storage
            if (response.ok) {
                const tokenString = data.jwt[0]
                const token = tokenString.split(';')[0].split('=')[1];
                // localStorage.setItem('token', token);
                // localStorage.setItem('email', logInEmail);
                setCookie('token', token, {
                    httpOnly: false,
                    // secure: true,
                    maxAge: 60 * 60 * 3,  // 3 hours
                })
                console.log('Login successful');
                console.log('Token:', token);
            } else {
                // Check if the data object has an error property
                if (data.error) {
                    // Extract the description from the error message
                    const descMatch = data.error.match(/desc = (.+)$/);
                    if (descMatch) {
                        console.log('Error description:', descMatch[1]);
                        // Optionally, you can use the description for user feedback
                        // For example, setting an error message state to display in the UI
                        setLogInServerError(descMatch[1]);
                    }
                }
                console.log('Sign up failed', data);
                setLoadingButton('');
                return;
            }
        } catch (error) {
            console.error('Error:', error);
            setLogInServerError('Internal server error');
            setLoadingButton('');
            return;
        }
        setLogInServerError('');
        setLoadingButton('');

        // Update zustand store
        setEmail(logInEmail);

        
        // Redirect to menu
        router.push('/menu');
    }

    const handleSignUp = async (e: React.FormEvent) => {
        e.preventDefault();

        
        // Verify all fields are filled
        if (renderErrorUsername === true || signUpEmail.length === 0 || signUpPassword.length === 0) {
            return;
        }
        setLoadingButton('Sign Up');

        const url = process.env.API_GATEWAY || 'https://xqtvgb1k-8080.asse.devtunnels.ms';
        const uri = '/rest/v1/user/signup';
        const endpoint = url + uri;
        try {
            // API call here - POST request to /rest/v1/user/signup, with email, password, and username in body
            const response = await fetch(endpoint, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    username: username,
                    email: signUpEmail,
                    password: signUpPassword,
                }),
            });
            const data = await response.json();
            if (response.ok) {
                console.log('Sign up successful', data);
            } else {
                // Check if the data object has an error property
                if (data.error) {
                    // Extract the description from the error message
                    const descMatch = data.error.match(/desc = (.+)$/);
                    if (descMatch) {
                        console.log('Error description:', descMatch[1]);
                        // Optionally, you can use the description for user feedback
                        // For example, setting an error message state to display in the UI
                        setSignUpServerError(descMatch[1]);
                    }
                }
                console.log('Sign up failed', data);
                setLoadingButton('');
                return;
            }
        } catch (error) {
            setSignUpServerError('Internal server error');
            setLoadingButton('');
            console.log('Failed to fetch: ', error);
            return;
        }
        setSignUpServerError('');
        setLoadingButton('');
        
        // Redirect to sent email page with email at query
        router.push('/sent-email?email=' + signUpEmail);
    }

    // Check for username format, always run on change. Length must longer than 3, must contain both alphabet and number, cannot have special characters
    useEffect(() => {
        if (username.length === 0) {
            setRenderErrorUsername(false);
        }
        else if (username.length < 3 || !username.match(/[a-zA-Z]/) || !username.match(/[0-9]/) || username.match(/[^a-zA-Z0-9]/)) {
            setRenderErrorUsername(true);
        }
        else {
            setRenderErrorUsername(false);
        }
    }, [username]);


    return (
        <div className='grid grid-cols-2 upper primary-background h-screen justify-items-center pt-[5rem] text-xl font-bold'>
            <div className='flex flex-col items-center'>
                <div className='text-5xl font-extrabold italic'>Log In</div>
                <form onSubmit={handleLogin} className='grid grid-cols-1 gap-[4rem] mt-[3.5rem]'>
                    <div>
                        <div className='indent-[1rem] mb-[1.25rem]'>Email</div>
                        <input type="email" placeholder="example@mail.com" value={logInEmail} onChange={(e) => setLogInEmail(e.target.value)} />
                    </div>
                    <div>
                        <div className='indent-[1rem] mb-[1.25rem]'>Password</div>
                        <input type="password" placeholder="Password" value={logInPassword} onChange={(e) => setLogInPassword(e.target.value)} />
                    </div>
                    <div className='flex flex-col'>
                        <button type="submit" className='bg-secondary w-fit rounded-lg self-center p-[0.35rem]'>{loadingButton === 'Login'? 'Loading...' : 'Log In'}</button>
                        {logInServerError !== ''? <p className='text-tertiary text-xs h-0 bg-primary overflow-visible whitespace-nowrap w-[15rem]'>{logInServerError}</p> : null}
                    </div>
                </form>
            </div>
            <div className='flex flex-col items-center'>
                <div className='text-5xl font-extrabold italic'>Sign Up</div>
                <form onSubmit={handleSignUp} className='grid grid-cols-1 gap-[2.3rem] mt-[3.5rem]'>
                    <div>
                        <div className='indent-[1rem] mb-[1.25rem]'>Username</div>
                        <input type="text" placeholder="Shine123" value={username} onChange={(e) => setUsername(e.target.value)} />
                        {renderErrorUsername? <p className='text-tertiary text-xs h-0 overflow-visible whitespace-nowrap w-[15rem]'>Must longer or equal than 3, contain both alphabet and number</p> : null}
                    </div>
                    <div>
                        <div className='indent-[1rem] mb-[1.25rem]'>Email</div>
                        <input type="email" placeholder="example@mail.com" value={signUpEmail} onChange={(e) => setSignUpEmail(e.target.value)} />
                    </div>
                    <div>
                        <div className='indent-[1rem] mb-[1.25rem]'>Password</div>
                        <input type="password" placeholder="Password" value={signUpPassword} onChange={(e) => setSignUpPassword(e.target.value)} />
                    </div>
                    <div className='flex flex-col'>
                    <button type="submit" className='bg-secondary w-fit rounded-lg self-center p-[0.35rem]'>{loadingButton === 'Sign Up'? 'Loading...' : 'Sign Up'}</button>
                    {signUpServerError !== ''? <p className='text-tertiary text-xs h-0 bg-primary overflow-visible whitespace-nowrap w-[15rem]'>{signUpServerError}</p> : null}
                    </div>
                </form>
            </div>
        </div>
    )
}
export default Login;