import { writable } from 'svelte/store';

export interface SymptomItem {
  slug: string;
  name: string;
  plants: number;
}

export interface FilterItem {
  type: string;
  term: string;
}

export const filters = writable<FilterItem[]>([]);