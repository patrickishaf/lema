import { Address } from "@/data/models";

export function parseAddress(address: Address): string {
  const { street, state, city, zipcode } = address;
  return `${street ?? ''}, ${state ?? ''} ${city ?? ''}, ${zipcode ?? ''}`;
}