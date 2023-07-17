import { getApiClient } from '@/stores/kotobako';
import axios from 'axios';

// GETリクエスト
export const getJSON = async (): Promise<any> => {
  const url = "http://localhost:8080/v1/kotonohas";
  const response = await axios.get(url);
  return response.data;
};

export const hello = async(): Promise<unknown> => {
  const client = getApiClient();
  const response = client.health({});
  return response;
}

export const getAllQuotes = async (): Promise<any> => {
  const url = "https://localhost:8080/v1/quotes";
  const response = await axios.get(url);
  return response.data;
}