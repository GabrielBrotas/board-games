import React, { useState } from "react";
import { MdClose } from "react-icons/md";

export const suggestedQuestions = [
  "Estamos em um lugar fechado ou aberto?",
  "Qual a cor predominante do local?",
  "Qual a temperatura do local?",
  "Qual o cheiro do local?",
  "Qual o som predominante do local?",
  "Esse é um local onde crianças frequentam?",
  "Você pode levar animais para esse local?",
  "Você levaria sua família para esse local?",
  "Esse local é seguro?",
  "Esse local é movimentado?",
  "Você pode comer nesse local?",
  "Esse local é público?",
  "Esse local é privado?",
  "Esse é um local de entretenimento ou trabalho?",
  "Quantas pessoas cabem nesse local?",
  "Quants janlas tem nesse local?",
  "Há muita movimentação de dinheiro nesse local?",
  "Na sua função, você interage com muitas pessoas?",
  "Você tem um uniforme para trabalhar nesse local?",
  "As pessoas utilizam uma tag para identificação nesse local?",
];

export const QuestionsModal = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);

  const toggleModal = () => setIsModalOpen(!isModalOpen);

  return (
    <div className="flex justify-center">
      <button
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
        onClick={toggleModal}
      >
        Visualizar Perguntas Sugeridas
      </button>

      <ModalComponent isOpen={isModalOpen} onClose={toggleModal}>
        <div className="flex flex-col gap-2">
          <h2 className="text-xl text-white font-bold mb-4">Sugestões</h2>
          <div className="overflow-y-auto">
            <ul className="list-disc list-inside px-4">
              {suggestedQuestions.map((question, index) => (
                <li
                  key={index}
                  className="text-white flex items-center gap-4 mb-4"
                >
                  {question}
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
