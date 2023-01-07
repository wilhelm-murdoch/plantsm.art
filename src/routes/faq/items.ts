export interface Item {
  question: string
  answer: string
  tags?: string[]
  open?: boolean
}

export const tags: object = {
  site: 'green',
  health: 'red'
}

export const items: Item[] = [{
  question: 'I noticed something is wrong or missing. What can I do?',
  answer: `The site is still under heavy development and the datasets that power it may not be 100% complete. If you notice a spelling error, incorrect image or if something is just not working for you, you are more than welcome to <a href="https://github.com/wilhelm-murdoch/plantsm.art/issues" title="Create an issue at Github.">let us know</a> about it or, if you're able to, <a href="/contribute" title="The contribute page.">contribute a fix</a> yourself.`,
  tags: [
    'site'
  ]
}, {
  question: 'I think my pet might have eaten a listed plant. What do I do?',
  answer: 'If you suspect your pet has ingested anything that may be unhealthy for them, immediately consult your local veterinarian\'s emergency hotline.',
  tags: [
    'health'
  ]
}, {
  question: 'Is this a complete dataset?',
  answer: 'No, this the dataset is not at all complete. However, it is constantly growing and improving. The ultimate goal is to cover as many common plants as we can for as many domesticated animal.',
}, {
  question: 'Does this site provide professional medical advice?',
  answer: 'No, this site should not be considered a reliable source of medical advice. It\'s sole purpose is to provide general guidance and to educate. If you suspect your pet has injested a poisonous substance, we urge you to contact your vet immediately.',
}]