import { QueueItmeLevel } from "yuumi-request"

export * from "./api-serve"

export interface RequestOption {
  level?: QueueItmeLevel;
  params?: { [key: string]: number|string };
  headers?: { [key: string]: string };
  cancelToken?: (cancel?: () => void) => void;
}
