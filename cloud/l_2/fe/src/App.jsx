import React, { useState, useEffect } from 'react';

const App = () => {
    const [time, setTime] = useState('');

    useEffect(() => {
        const updateTime = () => {
            const now = new Date();
            setTime(now.toLocaleTimeString());
        };

        const intervalId = setInterval(updateTime, 1000);
        updateTime(); // Initial call to set time immediately

        return () => clearInterval(intervalId); // Cleanup on unmount
    }, []);

    return (
        <div style={{ textAlign: 'center', marginTop: '20%' }}>
            <h1>Local Time</h1>
            <h2>{time}</h2>
        </div>
    );
};

export default App;
