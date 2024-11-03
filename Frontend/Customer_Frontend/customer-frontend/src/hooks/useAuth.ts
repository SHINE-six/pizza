// export const useAuth = () => {
// 	async function verifyCookieFromBackend() {
// 		// Send cookie to server
// 		try {
// 			const response = await fetch('https://xqtvgb1k-8080.asse.devtunnels.ms/rest/v1/verifyCookie', {
// 				method: 'GET',
// 				credentials: 'include',
// 			});
// 			console.log('Cookie sent to server');

// 			if (!response.ok) {
// 				// If the server response is not ok, throw an error
// 				throw new Error('Failed to verify cookie');
// 			}
			
// 			const data = await response.json();
// 			console.log('Cookie verified:', data);
// 			// If the server response is ok, return the data
// 			return data;

// 		} catch (error) {
// 			console.log('Failed to send cookie to server: ', error);
// 		}
// 	}

// 	return verifyCookieFromBackend;
// }

'use client';

import { useCallback, useEffect, useState } from 'react';
import { jwtDecode } from 'jwt-decode';
import { useUserStore } from '../../store/store';
import { getCookie, deleteCookie } from 'cookies-next';

type token = {
	email: string;
	exp: number;
	username: string;
	customerId: string;
}

export const useAuth = () => {
	const [isAuthenticated, setIsAuthenticated] = useState(false);
	const { setEmail, setUsername, setCustomerId } = useUserStore() as { setEmail: (value: string) => void, setUsername: (value: string) => void, setCustomerId: (value: string) => void };
	
	const checkExpiration = (exp: number) => {
		const currentTime = Date.now() / 1000;
		return exp > currentTime;
	}
	
	const decodeJWT = useCallback(() => {
		const token = getCookie('token');
		console.log('token:', token);
		// const token = localStorage.getItem('token');
		if (!token) {
			setIsAuthenticated(false);
			return;
		}
		
		const decodedToken = jwtDecode<token>(token);
		if (!checkExpiration(decodedToken.exp)) {
			setIsAuthenticated(false);
			deleteCookie('token');
			// localStorage.removeItem('token');
			return;
		}
		
		setIsAuthenticated(true);
		setEmail(jwtDecode<token>(token).email);
		setUsername(jwtDecode<token>(token).username);
		setCustomerId(jwtDecode<token>(token).customerId);

		console.log('decoded token:', decodedToken);
	}, [setEmail, setUsername, setCustomerId]);

	useEffect(() => {
		console.log('useAuth hook');
		decodeJWT();
	}, [decodeJWT]);

	return { isAuthenticated };
};