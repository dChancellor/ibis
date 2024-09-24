import { writable } from 'svelte/store';

export const activeModal = writable<HTMLDialogElement | null>(null);

export function openModal(dialog: HTMLDialogElement) {
    activeModal.set(dialog);
    dialog.showModal();
}

export function closeModal(dialog: HTMLDialogElement) {
    dialog.close();
    activeModal.set(null);
}
