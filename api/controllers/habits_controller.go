package controllers

// GetResumoPGM godoc
// @Summary Resumo do PGM
// @Description Obtém um resumo do PGM com base nos parâmetros fornecidos.
// @Tags Gestor de Transações
// @Accept  json
// @Produce  json
// @Param   items      query    int true  "Quantidade de Itens"
// @Param   page       query    int true  "Número da Página"
// @Param   data_inicio  query    string  true  "Data de Início"
// @Param   data_fim    query    string  true  "Data de Fim"
// @Param   estado    query    string  true  "Estado do cliente"
// @Param   projeto    query    string  true  "Projeto do cliente"
// @Param   documento    query    string  true  "Documento do cliente"
// @Success 200 {object}  models.Resumo
// @Router /resumo [get]
// func GetResumoPGM() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		_, err := funcs.ValidarToken(r)

// 		if err != nil {
// 			responses.HTTPError(w, "Erro ao validar token: "+err.Error(), 401)
// 			return
// 		}

// 		db := database.ConnDb()
// 		defer db.Close()

// 		//Definindo filtros
// 		var items = funcs.GetParamsInt(r.URL.Query(), "items", 1, 10)
// 		var page = funcs.GetParamsInt(r.URL.Query(), "page", 1, 1)
// 		projetos := r.URL.Query()["projeto"]
// 		estados := r.URL.Query()["estado"]
// 		documentos := r.URL.Query()["documento"]
// 		var offset = (page - 1) * items

// 		data_inicio, err := funcs.GetParamsDate(r.URL.Query().Get("data_inicio"))

// 		if err != nil {
// 			responses.HTTPError(w, "Erro: data de entrada em formato inválido", 400)
// 			return
// 		}

// 		data_fim, err := funcs.GetParamsDate(r.URL.Query().Get("data_fim"))

// 		if err != nil {
// 			responses.HTTPError(w, "Erro: data de entrada em formato inválido", 400)
// 			return
// 		}

// 		transacoes, clientes, recebimentos, creditos, err := funcs.ObterResumoV2(db, r.Context(), data_inicio, data_fim, estados, projetos, documentos, items, offset)

// 		if err != nil {
// 			responses.HTTPError(w, err.Error(), 500)
// 			return
// 		}

// 		//Criei esse mapeamento para diminuir o numero de consultas,
// 		//Anteriormente fiz uma consulta que busca tudo de uma vez e estava com
// 		//Problemas de Timeout no http
// 		clienteMap := make(map[string]models.Cliente)
// 		for _, c := range clientes {
// 			clienteMap[c.Documento] = c
// 		}

// 		recebimentoMap := make(map[string]models.Recebimento)
// 		for _, r := range recebimentos {
// 			key := r.Autorizacao + "-" + r.Nsu + "-" + r.Documento
// 			recebimentoMap[key] = r
// 		}

// 		creditoMap := make(map[string]models.Credito)
// 		for _, c := range creditos {
// 			key := c.Autorizacao + "-" + c.Nsu + "-" + c.Documento
// 			creditoMap[key] = c
// 		}

// 		var resumoSlice []models.Resumo

// 		for _, t := range transacoes {
// 			if cliente, ok := clienteMap[t.Documento]; ok {
// 				var resumo models.Resumo
// 				resumo.IdVenda = t.Id
// 				resumo.Nsu = t.Nsu
// 				resumo.Autorizacao = t.Autorizacao
// 				resumo.Documento = t.Documento
// 				resumo.ValorTransacao = t.ValorVenda
// 				resumo.Parcelas = t.Parcelas
// 				resumo.Bandeira = t.Bandeira
// 				resumo.DataTransacao = t.DataTransacao
// 				resumo.HoraTransacao = t.HoraTransacao
// 				resumo.Status = t.Status
// 				resumo.StatusDescricao = t.StatusDescricao

// 				resumo.LojistaID = cliente.Id
// 				resumo.Nome = cliente.Nome
// 				resumo.Documento = cliente.Documento
// 				resumo.Projeto = cliente.Projeto
// 				resumo.UF = cliente.UF

// 				recebimentoKey := t.Autorizacao + "-" + t.Nsu + "-" + t.Documento
// 				if recebimento, ok := recebimentoMap[recebimentoKey]; ok {
// 					resumo.TotalRecebimento = recebimento.TotalRecebimentos
// 					resumo.TaxaRecebimento = recebimento.Taxa
// 					resumo.PorcentagemSplit = recebimento.PorcentagemSplit
// 					resumo.PagadorSplit = recebimento.RegraSplit
// 					resumo.Modalidade = recebimento.Modalidade
// 				}

// 				if credito, ok := creditoMap[recebimentoKey]; ok {
// 					resumo.CreditoGerado = credito.Valor
// 					resumo.TaxaCredito = credito.Taxa
// 				}

// 				resumoSlice = append(resumoSlice, resumo)
// 			}
// 		}

// 		total_registros, err := funcs.ObterTotalRegistros(db, r.Context(), data_inicio, data_fim, estados, projetos, documentos)

// 		if err != nil {
// 			responses.HTTPError(w, err.Error(), 500)
// 			return
// 		}

// 		responses.HTTPSuccessPaginate(w, map[string]interface{}{"resumo": resumoSlice}, page, items, total_registros)

// 	}
// }
