import React from "react";

interface GameSetupProps {
  imposterChances: {
    one: number;
    two: number;
    three: number;
  };
  handleImposterChange: (key: string, value: string) => void;
  category: string;
  setCategory: (category: string) => void;
  difficulty: string;
  setDifficulty: (difficulty: string) => void;
  resetPoints: () => void;
}

export const GameSetup = ({
  imposterChances,
  handleImposterChange,
  category,
  setCategory,
  difficulty,
  setDifficulty,
  resetPoints,
}: GameSetupProps) => {
  return (
    <div className="bg-gray-900 text-white p-8 rounded-lg shadow-lg w-full">
      <h1 className="text-xl font-bold mb-6 text-center">Setup Game</h1>
      <div className="flex justify-between items-center mb-4 flex-wrap">
        <label htmlFor="one-imposter" className="flex-1 text-sm font-semibold">
          1 Imposter:
        </label>
        <input
          type="range"
          id="one-imposter"
          value={imposterChances.one}
          onChange={(e) => handleImposterChange("one", e.target.value)}
          className="flex-2 range range-primary"
          min="0"
          max="100"
        />
        <span className="ml-4 w-12 text-center">{imposterChances.one}%</span>
      </div>
      <div className="flex justify-between items-center mb-4">
        <label htmlFor="two-imposters" className="flex-1 text-sm font-semibold">
          2 Imposters:
        </label>
        <input
          type="range"
          id="two-imposters"
          value={imposterChances.two}
          onChange={(e) => handleImposterChange("two", e.target.value)}
          className="flex-2 range range-primary"
          min="0"
          max="100"
        />
        <span className="ml-4 w-12 text-center">{imposterChances.two}%</span>
      </div>
      <div className="flex justify-between items-center mb-6">
        <label
          htmlFor="three-imposters"
          className="flex-1 text-sm font-semibold"
        >
          3 Imposters:
        </label>
        <input
          type="range"
          id="three-imposters"
          value={imposterChances.three}
          onChange={(e) => handleImposterChange("three", e.target.value)}
          className="flex-2 range range-primary"
          min="0"
          max="100"
        />
        <span className="ml-4 w-12 text-center">{imposterChances.three}%</span>
      </div>
      <div>
        <label htmlFor="category" className="text-sm font-semibold">
          Categoria:
        </label>
        <input
          type="text"
          id="category"
          className="bg-gray-700 text-white p-2 rounded w-full"
          value={category}
          onChange={(e) => setCategory(e.target.value)}
        />
      </div>
      <div className="mt-2">
        <label htmlFor="difficulty" className="text-sm font-semibold">
          Dificuldade:
        </label>
        <input
          type="text"
          id="difficulty"
          className="bg-gray-700 text-white p-2 rounded w-full"
          value={difficulty}
          onChange={(e) => setDifficulty(e.target.value)}
        />
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
