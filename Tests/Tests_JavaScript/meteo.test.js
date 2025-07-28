import { it, expect, vi, beforeEach } from 'vitest'
import { displayMeteo, fetchWeather } from './meteo'
import fetch from 'node-fetch'

// Mock de node-fetch
vi.mock('node-fetch', () => ({
  default: vi.fn(),
}))

// Nettoyer les mocks avant chaque test
beforeEach(() => {
  vi.clearAllMocks()
})


it('Test de la fonction displayMeteo avec une ville inexistante', async () => {
  // Simulation de la réponse API pour ville inexistante
  const emptyResponse = {
    generationtime_ms: 0.059723854
  }

  fetch.mockResolvedValueOnce({
    ok: true,
    json: vi.fn().mockResolvedValue(emptyResponse),  // ✅ Mock function propre
  })

  // On s'attend à ce que la fonction lance un TypeError
  await expect(displayMeteo("VilleInexistante123")).rejects.toThrow()
  
  // Vérifier que l'API de géocodage a bien été appelée
  expect(fetch).toHaveBeenCalledWith(
    'https://geocoding-api.open-meteo.com/v1/search?name=VilleInexistante123'
  )

})

 

 


