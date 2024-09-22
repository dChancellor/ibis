import {writable} from 'svelte/store'
import { main } from '../../wailsjs/go/models'

export const activeSkill = writable<main.Skill>()

export const skills = writable<main.Skill[]>([])