'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';

export default function Home() {
  const router = useRouter();
  const [signUpServerError, setSignUpServerError] = useState('');

  async function verifyLoginWithServer(email: string) {
    const queryParams = new URLSearchParams({ email }).toString();
    const url = process.env.API_GATEWAY || 'https://xqtvgb1k-8080.asse.devtunnels.ms';
    const uri = `/rest/v1/staff/verify_email?${queryParams}`;
    const endpoint = url + uri;
    try {
      const response = await fetch(endpoint, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

        const data = await response.json();
        console.log(data);
        // Check if the data object has an error property
        if (data.error) {
          // Extract the description from the error message
          const descMatch = data.error.match(/desc = (.+)$/);
          if (descMatch) {
              console.log('Error description:', descMatch[1]);
              // For example, setting an error message state to display in the UI
              setSignUpServerError(descMatch[1]);
          }
        } else {
          console.log('Sign up successful', data);
          // Redirect to pickOrder page with staffID and staffName at query
          router.push('/order?id=' + data.data.staffID + '&name=' + data.data.staffName);
        }

    } catch (error) {
      console.error(error);
    }
  }


  function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();

    const form = event.currentTarget;
    const email = form.email.value;
    console.log(email)
    verifyLoginWithServer(email);
  }

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <form onSubmit={handleSubmit} className="flex flex-col gap-4">
        <label htmlFor="email">Email</label>
        <input type="email" name="email" id="email" />
        <button type="submit">Login</button>
      </form>
      {signUpServerError && <p className="text-red-500">{signUpServerError}</p>}
    </main>
  );
}
