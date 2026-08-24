import { describe, it, expect } from 'vitest'
import { render, screen } from '@testing-library/react'
import App from './App'

describe('App', () => {
  it('renders the Renova title', () => {
    render(<App />)
    const titleElement = screen.getByRole('heading', { name: /Renova/i })
    expect(titleElement).toBeInTheDocument()
  })

  it('renders the description', () => {
    render(<App />)
    const descriptionElement = screen.getByText(/Aplicativo de gestão de renovação de documentos/i)
    expect(descriptionElement).toBeInTheDocument()
  })
})