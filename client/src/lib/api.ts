import axios from "axios";

const baseAPI = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
});

type CreateUserResponse = {
  id: string;
  name: string;
};

const createUserOrLogin = async (name: string): Promise<CreateUserResponse> => {
  const response = await baseAPI.post(`/login`, {
    name,
  });

  return response.data;
};

type ImposterGameStatusResponse = {
  gameStarted: boolean;
  word: string;
  inGame: boolean;
};

const getImposterGameStatus = async (
  userID: string
): Promise<ImposterGameStatusResponse> => {
  const response = await baseAPI.get(`/games/impostor/status?u=${userID}`);
  return response.data;
};

type SpyfallGameStatusResponse = {
  gameStarted: boolean;
  inGame: boolean;
  role: string;
  location: string;
};

const getSpyfallGameStatus = async (
  userID: string
): Promise<SpyfallGameStatusResponse> => {
  const response = await baseAPI.get(`/games/spyfall/status?u=${userID}`);
  return response.data;
};



export const api = {
  createUserOrLogin,
  getImposterGameStatus,
  getSpyfallGameStatus,
};
