'use client';

import { useState, useEffect } from 'react';
import { LuMoon } from "react-icons/lu";
import { MdOutlineWbSunny } from "react-icons/md";

const ThemeToggleButton: React.FC = () => {
    // Initialize theme state without immediately applying the theme
    const [theme, setTheme] = useState<string>('light');

    useEffect(() => {
        // Apply theme from localStorage or default to 'light' once the component mounts
        const storedTheme = localStorage.getItem('theme') || 'light';
        setTheme(storedTheme);
    }, []);

    useEffect(() => {
        document.documentElement.setAttribute('data-theme', theme);
        localStorage.setItem('theme', theme);
    }, [theme]);

    const toggleTheme = () => {
        setTheme((prevTheme) => (prevTheme === 'light' ? 'dark' : 'light'));
    };

    return (
        <button onClick={toggleTheme}>
            {theme === 'light' ? <MdOutlineWbSunny /> : <LuMoon />}
        </button>
    )
};

export default ThemeToggleButton;
