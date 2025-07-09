import { ref, type Ref } from "vue";
export type ApiRequest = () => Promise<void>;
export type ApiOptions = RequestInit;

export default function useApi<T>(url: RequestInfo, options: ApiOptions){
  // ToDo: Error Handling
  // ToDo: API Security, add Headers for currently logged in User
  options.headers = {
    "UserSession": "IchBinEinBier"
  }

  let response: Ref<T | undefined>= ref();
  let request: ApiRequest = async () => {
    let res = await fetch(url, options);
    if (!res.ok){
      throw new Error(`Response status: ${res.status}`);
    }

    let data = await res.json();
    response.value = data;
  }
  return {response, request};
}

export function getApi<T>(url: RequestInfo){
  return useApi<T>(url, {
    method: "GET"
  })
}

export function putApi<T>(url: RequestInfo, data?: RequestInit["body"]){
  return useApi<T>(url, {
    method: "PUT",
    body: data
  })
}

export function postApi<T>(url: RequestInfo, data?: RequestInit["body"]){
  return useApi<T>(url, {
    method: "POST",
    body: data
  })
}

export function deleteApi<T>(url: RequestInfo, data?: RequestInit["body"]){
  return useApi<T>(url, {
    method: "DELETE",
    body: data
  })
}