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
  const client = getApiClient();
  const response = client.listQuotes({});
  return response;
}

export const getQuote = async (id:string): Promise<any> => {
  const client = getApiClient();
  const response = client.getQuote({quoteId : String(id)});
  return response;
}