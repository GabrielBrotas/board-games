import React, { useState } from "react";
import { MdClose } from "react-icons/md";

export const availableLocations = [
  {
    location: "Hospital",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/hospital.jpg",
  },
  {
    location: "Estação Espacial",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/space-station.jpg",
  },
  {
    location: "Supermercado",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/supermarket.jpg",
  },
  {
    location: "Submarino",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/submarine.jpg",
  },
  {
    location: "Banco",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/bank.jpg",
  },
  {
    location: "Escola",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/school.jpg",
  },
  {
    location: "Circo",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/circus.jpg",
  },
  {
    location: "Restaurante",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/restaurant.jpg",
  },
  {
    location: "Teatro",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/theater.jpg",
  },
  {
    location: "Aeroporto",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/airport.jpg",
  },
  {
    location: "Zoológico",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/zoo.jpg",
  },
  {
    location: "Cassino",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/casino.jpg",
  },
  {
    location: "Navio Cruzeiro",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/cruise-ship.jpg",
  },
  {
    location: "Parque de Diversões",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/amusement-park.jpg",
  },
  {
    location: "Museu",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/museum.jpg",
  },
  {
    location: "Estúdio de TV",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/tv-studio.jpg",
  },
  {
    location: "Avião",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/airplane.jpg",
  },
  {
    location: "Praia",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/beach.jpg",
  },
  {
    location: "Cinema",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/movie-theater.jpg",
  },
  {
    location: "Base Militar",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/military-base.jpg",
  },
  {
    location: "Spa",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/spa.jpg",
  },
  {
    location: "Trem",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/train.jpg",
  },
  {
    location: "Delegacia",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/police-station.jpg",
  },
  {
    location: "Oficina",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/car-workshop.jpg",
  },
  {
    location: "Estádio de Futebol",
    image: "https://raw.githubusercontent.com/GabrielBrotas/board-games/main/images/spyfall/football-stadium.jpg",
  },
];

export const LocationModal = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);

  const toggleModal = () => setIsModalOpen(!isModalOpen);

  return (
    <div className="flex justify-center">
      <button
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
        onClick={toggleModal}
      >
        Visualizar Localizações
      </button>

      <ModalComponent isOpen={isModalOpen} onClose={toggleModal}>
        <div className="flex flex-col gap-2">
          <h2 className="text-xl text-white font-bold mb-4">Localizações</h2>
          <div className="overflow-y-auto">
            <ul className="list-disc list-inside px-4">
              {availableLocations.map((location, index) => (
                <li
                  key={index}
                  className="text-white flex items-center gap-4 mb-4"
                >
                  {location.image && (
                    <img
                      src={location.image}
                      alt={location.location}
                      className="w-32 h-32 object-cover"
                    />
                  )}
                  {location.location}
                </li>
              ))}
            </ul>
          </div>
        </div>
      </ModalComponent>
    </div>
  );
};

interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
  children: React.ReactNode;
}

const ModalComponent = ({ isOpen, onClose, children }: ModalProps) => {
  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 z-50 flex justify-center items-center p-4">
      <button
        onClick={onClose}
        className="absolute top-2 right-2 text-white text-2xl rounded focus:outline-none"
      >
        <MdClose size={25} />
      </button>
      <div className="relative bg-gray-900 p-4 max-w-lg w-full md:max-w-md mx-auto rounded overflow-auto max-h-[80vh]">
        {children}
      </div>
    </div>
  );
};
