import React, { useState } from 'react';

const App = () => {
    const [input, setInput] = useState('');

    const handleButtonClick = (value) => {
        setInput((prev) => prev + value);
    };

    const handleClear = () => {
        setInput('');
    };

    const handleCalculate = () => {
        try {
            setInput(eval(input).toString());
        } catch {
            setInput('Error');
        }
    };

    return (
        <div className="calculator">
            <input
                type="text"
                value={input}
                readOnly
                className="display"
            />
            <div className="buttons">
                {['7', '8', '9', '/'].map((item) => (
                    <button
                        key={item}
                        onClick={() => handleButtonClick(item)}
                        className="button"
                    >
                        {item}
                    </button>
                ))}
                {['4', '5', '6', '*'].map((item) => (
                    <button
                        key={item}
                        onClick={() => handleButtonClick(item)}
                        className="button"
                    >
                        {item}
                    </button>
                ))}
                {['1', '2', '3', '-'].map((item) => (
                    <button
                        key={item}
                        onClick={() => handleButtonClick(item)}
                        className="button"
                    >
                        {item}
                    </button>
                ))}
                {['0', '.', '=', '+'].map((item) => (
                    <button
                        key={item}
                        onClick={item === '=' ? handleCalculate : () => handleButtonClick(item)}
                        className={`button ${item === '=' ? 'operator' : ''}`}
                    >
                        {item}
                    </button>
                ))}
                <button
                    onClick={handleClear}
                    className="button clear"
                >
                    Clear
                </button>
            </div>
        </div>
    );
};

export default App;
