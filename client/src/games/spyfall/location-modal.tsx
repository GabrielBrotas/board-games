import React, { useState } from "react";
import { MdClose } from "react-icons/md";

const locations = [
  {
    Location: "Hospital",
    slug: "hospital",
    image: "/images/spyfall/hospital.jpg",
  },
  {
    Location: "Estação Espacial",
    slug: "estacao-espacial",
  },
  {
    Location: "Supermercado",
    slug: "supermercado",
  },
  {
    Location: "Submarino",
    slug: "submarino",
  },
  {
    Location: "Banco",
    slug: "banco",
  },
  {
    Location: "Escola",
    slug: "escola",
  },
  {
    Location: "Circo",
    slug: "circo",
  },
  {
    Location: "Restaurante",
    slug: "restaurante",
  },
  {
    Location: "Teatro",
    slug: "teatro",
  },
  {
    Location: "Aeroporto",
    slug: "aeroporto",
  },
  {
    Location: "Zoológico",
    slug: "zoo",
  },
  {
    Location: "Cassino",
    slug: "cassino",
  },
  {
    Location: "Navio Cruzeiro",
    slug: "navio-cruzeiro",
  },
  {
    Location: "Parque de Diversões",
    slug: "parque-diversoes",
  },
  {
    Location: "Museu",
    slug: "museu",
  },
  {
    Location: "Estúdio de TV",
    slug: "estudio-tv",
  },
];

const LocationModal = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);

  const toggleModal = () => setIsModalOpen(!isModalOpen);

  return (
    <div className="flex justify-center">
      <button
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded m-4"
        onClick={toggleModal}
      >
        Visualizar Localizações
      </button>

      <ModalComponent isOpen={isModalOpen} onClose={toggleModal}>
        <div className="flex flex-col gap-2">
          <h2 className="text-xl text-white font-bold mb-4">Localizações</h2>
          <div className="overflow-y-auto">
            <ul className="list-disc list-inside px-4">
              {locations.map((location, index) => (
                <li
                  key={index}
                  className="text-white flex items-center gap-4 mb-4"
                >
                  <img
                    src={"/images/spyfall/hospital.jpg"}
                    alt={location.Location}
                    className="w-32 h-32 object-cover"
                  />
                  {location.Location}
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

export default LocationModal;
