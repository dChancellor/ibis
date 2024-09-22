import {writable} from 'svelte/store'
import { main } from '../../wailsjs/go/models'

export const constants = writable<main.MessagesStruct>()
