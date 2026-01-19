const iconMap: Record<string, string> = {
  '0': '📄',
  '1': '📁',
  '2': '☎️',
  '3': '❌',
  '4': '📦',
  '5': '📦',
  '6': '📦',
  '7': '🔍',
  '8': '💻',
  '9': '📦',
  '+': '🔄',
  'g': '🖼',
  'I': '🖼',
  'T': '💻',
  'h': '🌐',
  'i': 'ℹ️',
  's': '🔊',
};

export function getIcon(type: string): string {
  if (type in iconMap) {
    return iconMap[type];
  }
  return '❓';
}

const clickableTypes = [
  '0',
  '1',
  '7',
  'g',
  'I',
]

export function isClickableType(type: string): boolean {
  return clickableTypes.includes(type);
}

export function isInfoType(type: string): boolean {
  return type === 'i' || type === '3';
}