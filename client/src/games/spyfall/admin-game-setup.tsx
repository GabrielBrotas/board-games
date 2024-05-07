import React from "react";

interface GameSetupProps {
  spiesChances: {
    one: number;
    two: number;
    three: number;
  };
  handleSpiesChange: (key: string, value: string) => void;
  resetPoints: () => void;
}

export const GameSetup = ({
  resetPoints,
  handleSpiesChange,
  spiesChances,
}: GameSetupProps) => {
  return (
    <div className="bg-gray-900 text-white p-4 rounded-lg shadow-lg max-w-lg w-full">
      <h1 className="text-xl font-bold mb-6 text-center">Game Setup</h1>

      <div className="flex justify-between items-center mb-4">
        <label htmlFor="one-spy" className="flex-1 text-sm font-semibold">
          1 Spy:
        </label>
        <input
          type="range"
          id="one-spy"
          value={spiesChances.one}
          onChange={(e) => handleSpiesChange("one", e.target.value)}
          className="flex-2 range range-primary"
          min="0"
          max="100"
        />
        <span className="ml-4 w-12 text-center">{spiesChances.one}%</span>
      </div>
      <div className="flex justify-between items-center mb-4">
        <label htmlFor="two-spies" className="flex-1 text-sm font-semibold">
          2 Spies:
        </label>
        <input
          type="range"
          id="two-spies"
          value={spiesChances.two}
          onChange={(e) => handleSpiesChange("two", e.target.value)}
          className="flex-2 range range-primary"
          min="0"
          max="100"
        />
        <span className="ml-4 w-12 text-center">{spiesChances.two}%</span>
      </div>
      <div className="flex justify-between items-center mb-6">
        <label htmlFor="three-spies" className="flex-1 text-sm font-semibold">
          3 Spies:
        </label>
        <input
          type="range"
          id="three-spies"
          value={spiesChances.three}
          onChange={(e) => handleSpiesChange("three", e.target.value)}
          className="flex-2 range range-primary"
          min="0"
          max="100"
        />
        <span className="ml-4 w-12 text-center">{spiesChances.three}%</span>
      </div>

      <button
        className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded mt-4 w-full"
        onClick={resetPoints}
      >
        Resetar pontos
      </button>
    </div>
  );
};
