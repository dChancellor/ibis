import {writable} from 'svelte/store'
import { application } from '../../wailsjs/go/models'

export const constants = writable<application.MessagesStruct>()
