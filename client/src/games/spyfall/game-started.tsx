import React, { useEffect, useState } from "react";
import { availableLocations } from "./location-modal";

interface GameStartedProps {
  inGame: boolean;
  role: string;
  location: string;
}

export const SPY_ROLE = "spy";

export const GameStarted = ({
  inGame,
  role,
  location,
}: GameStartedProps) => {
  const [locationImg, setLocationImg] = useState(<></>);

  useEffect(() => {
    if (location) {
      const locationFound = availableLocations.find(
        (loc) => loc.location === location
      );
      if (locationFound?.image) {
        setLocationImg(
          <img
            src={locationFound.image}
            alt={locationFound.location}
            className="w-32 h-32 object-cover"
          />
        );
      }
    }
  }, [location]);
  return (
    <>
      <div className="bg-gray-900 text-white p-8 rounded-lg shadow-lg max-w-sm">
        <h1 className="text-lg font-bold mb-4 text-center">Quem é o Spy?</h1>
        {inGame ? (
          role == SPY_ROLE ? (
            <div className="">
              <p className="mb-4 text-center">Você é o Spy!</p>
            </div>
          ) : (
            <div className="flex flex-col items-start gap-4">
              <p className="">Localização: {location}</p>
              <div className="flex justify-center w-full">{locationImg}</div>
              <p className="">Função: {role}</p>
            </div>
          )
        ) : (
          <p className="mb-4 text-center">
            Esperando jogadores terminar a partida...
          </p>
        )}
      </div>
    </>
  );
};
