export function parseFeatureBoolean(featureValue?: string): boolean {
  if (!featureValue) {
    return false;
  }
  const lower = featureValue.trim().toLowerCase();
  return lower === 'true' || lower === '1';
}
