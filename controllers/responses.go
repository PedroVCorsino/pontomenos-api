package controllers

// RegistroPontoResponse define a estrutura de resposta para registros de ponto.
type RegistroPontoResponse struct {
    ID        uint      `json:"id"`
    UsuarioID uint      `json:"usuario_id"`
    DataHora  string    `json:"data_hora"`
    TipoPonto string    `json:"tipo_ponto"`
}
