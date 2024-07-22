package models

import (
	"time"

	"github.com/uptrace/bun"
)

type (
	FaixaParcela struct {
		FinalParc   int     `json:"finalParc"`
		RateValue   float64 `json:"rateValue"`
		InitialParc int     `json:"initialParc"`
	}

	Taxa struct {
		RateValue     float64         `json:"rateValue"`
		ProductTrans  int             `json:"productTrans"`
		FaixaParcelas *[]FaixaParcela `json:"faixaParcelas,omitempty"`
	}

	Taxas struct {
		Taxas     []Taxa `json:"taxas"`
		CardBrand string `json:"cardBrand"`
	}

	DashboardPlanos struct {
		bun.BaseModel `bun:"table:credenciamento.dashboard_planos,alias:dp"`

		ID                  int       `bun:"id,pk,autoincrement"`
		Nome                string    `bun:"nome,notnull"`
		TipoContratacao     string    `bun:"tipo_contratacao,notnull"`
		ValorContratacao    float64   `bun:"valor_contratacao,notnull"`
		TaxaDebito          float64   `bun:"taxa_debito,notnull"`
		TaxaCredito         float64   `bun:"taxa_credito,notnull"`
		TaxaParcelado       float64   `bun:"taxa_parcelado,notnull"`
		ProdutoID           int       `bun:"produto_id,notnull"`
		Adquirencia         string    `bun:"adquirencia,notnull"`
		DataCadastro        time.Time `bun:"data_cadastro,nullzero"`
		DataUpdate          time.Time `bun:"data_update,nullzero"`
		JSONPrazos          string    `bun:"json_prazos,nullzero"`
		JSONTaxasBandeiras  []Taxas   `bun:"json_taxasbandeiras,nullzero"`
		TipoProjeto         string    `bun:"tipo_projeto,nullzero,default:null"`
		InicioVigencia      time.Time `bun:"inicio_vigencia,nullzero"`
		FinalVigencia       time.Time `bun:"final_vigencia,nullzero"`
		ProjetoID           int       `bun:"projeto_id,nullzero"`
		Tipo                string    `bun:"tipo,nullzero"`
		Antecipacao         bool      `bun:"antecipacao,default:false"`
		ValorAntecipacao    float64   `bun:"valor_antecipacao,nullzero"`
		IDParent            int       `bun:"id_parent,nullzero"`
		Cargos              string    `bun:"cargos,nullzero"`
		Ordem               int       `bun:"ordem,nullzero"`
		Ativo               bool      `bun:"ativo,notnull,default:true"`
		Simulacao           bool      `bun:"simulacao,notnull,default:true"`
		SplitFavorecidos    string    `bun:"split_favorecidos,notnull,default:'[]'"`
		SplitCredenciamento bool      `bun:"split_credenciamento,notnull,default:false"`
		SplitAtivo          bool      `bun:"split_ativo,notnull,default:false"`
		PlanoTipoID         int       `bun:"plano_tipo_id,nullzero"`
		P18                 bool      `bun:"p18,notnull,default:false"`
		Especial            bool      `bun:"especial,notnull,default:false"`
		Grupo               string    `bun:"grupo,nullzero"`
		Regiao              string    `bun:"regiao,nullzero"`
	}

	AuditoriaPlanos struct {
		bun.BaseModel `bun:"table:credenciamento.auditoria_planos,alias:ap"`

		ID                   int       `bun:"id,pk,autoincrement"`
		LojistaID            int       `bun:"lojista_id,nullzero"`
		PlanoID              int       `bun:"plano_id,nullzero"`
		DataInicioUtilizacao time.Time `bun:"data_inicio_utilizacao,nullzero"`
		DataFimUtilizacao    time.Time `bun:"data_fim_utilizacao,nullzero"`
		ResponsavelID        int       `bun:"responsavel_id,nullzero"`
		Sistema              string    `bun:"sistema,nullzero"`
		CreatedAt            time.Time `bun:"created_at,notnull,default:now()"`
	}

	AuditoriaTaxas struct {
		bun.BaseModel `bun:"table:credenciamento.auditoria_taxas,alias:ax"`

		ID            int       `bun:"id,pk,autoincrement"`
		IdPlano       int       `bun:"id_plano"`
		TaxasAntigas  []Taxas   `bun:"taxas_antigas"`
		TaxasNovas    []Taxas   `bun:"taxas_novas"`
		DataAlteracao time.Time `bun:"data_alteracao"`
		Sistema       string    `bun:"sistema"`
	}

	AuditoriaIndividualizacaoPlanos struct {
		bun.BaseModel `bun:"table:credenciamento.auditoria_individualizacao_planos,alias:aip"`

		ID            int       `bun:"id,pk,autoincrement"`
		IdPlanoOrigem int       `bun:"id_plano_origem"`
		IdPlanoCriado int       `bun:"id_plano_criado"`
		IdLojista     int       `bun:"id_lojista"`
		CreatedAt     time.Time `bun:"created_at,default:now()"`
	}

	LojistaPlano struct {
		bun.BaseModel  `bun:"table:credenciamento.dashboard_lojista,alias:dl"`
		ID             int    `bun:"id,pk,autoincrement"`
		Nome           string `bun:"nome"`
		Documento      string `bun:"documento"`
		PlanoID        int    `bun:"plano_id"`
		PlanoVirtualId int    `bun:"plano_virtual_id"`
		Plano18xId     int    `bun:"plano_18x_id"`
	}
)
