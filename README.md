## 🚀 Desafio: Corrida de APIs de CEP (Go)

Neste desafio, o objetivo é fazer chamadas simultâneas para duas APIs de CEP e exibir a resposta da que retornar primeiro.

### APIs envolvidas

- [BrasilAPI](https://brasilapi.com.br/api/cep/v1/{cep})
- [ViaCEP](http://viacep.com.br/ws/{cep}/json/)

---

### ✅ Requisitos

- Fazer as duas requisições simultaneamente.
- Exibir a resposta da API que retornar mais rápido.
- Ignorar a resposta mais lenta.
- Exibir no terminal os dados do endereço e qual API respondeu.
- Limitar o tempo de resposta em 1 segundo.
- Se nenhuma API responder a tempo, mostrar erro de timeout.

---

### 🛠️ Etapas para implementação

1. **Preparar o ambiente**
   - Criar um projeto Go (`go mod init`).
   - Importar os pacotes necessários

2. **Definir o input (CEP)**
   - Declarar o CEP fixo ou aceitar como argumento/entrada.

3. **Criar structs para os dados**
   - Mapear os campos de resposta esperados de cada API usando structs.

4. **Criar funções de request**
   - Função para consumir a BrasilAPI.
   - Função para consumir a ViaCEP.
   - Ambas devem receber `context.Context`.

5. **Configurar canal de resposta**
   - Criar um `chan` para receber a resposta da API que retornar primeiro.
   - Executar cada requisição em goroutines separadas.

6. **Controlar timeout com context**
   - Criar `context.WithTimeout` de 1 segundo.
   - Passar o contexto para as funções de request.

7. **Capturar e exibir a primeira resposta**
   - Ouvir o canal.
   - Cancelar o contexto assim que a primeira resposta chegar.
   - Exibir os dados no terminal (endereço + nome da API).

8. **Tratar caso de timeout**
   - Se nenhuma resposta chegar em até 1 segundo, exibir mensagem de timeout.

---
