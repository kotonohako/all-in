import { KotobakoService } from "@/generated/buf/kotobako/v1/kotobako_connect";
import { createPromiseClient } from "@bufbuild/connect";
import { createConnectTransport } from "@bufbuild/connect-web";

export const getApiClient = () => {
    const transport = createConnectTransport( { 
        baseUrl: "http://localhost:8080/api/"
    });
    const client = createPromiseClient(KotobakoService, transport);

    return client;
}