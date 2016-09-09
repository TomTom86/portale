package controllers

import (
	"fmt"
	"html/template"
	_ "portale/models" //importo sempre models per evitare errori
	"portale/modules/automezzi/models"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	_ "github.com/mattn/go-sqlite3" // importo sempre squid per evitare errori
)

var (
	appcfgdomainname = beego.AppConfig.String("appcfgdomainname")
)

//AutomezziController è il controller beego
type AutomezziController struct {
	beego.Controller
}

type automezzo struct {
	Targa                string `form:"targa" valid:"Required"`
	DataInFlotta         string `form:"dataInFlotta" valid:"Required"`
	DataFineFlotta       string `form:"dataFineFlotta"`
	TipoVeicolo          int    `form:"tipoveicolo" valid:"Required"`
	Settore              int    `form:"settore" valid:"Required"`
	Condizione           int    `form:"condizione" valid:"Required"`
	Impiego              int    `form:"impiego" valid:"Required"`
	Conducente           int    `form:"conducente" valid:"Required"`
	Note                 string `form:"note"`
	AnnoImmatricolazione int    `form:"annoImmatricolazione" valid:"Required"`
	NLibretto            string `form:"numeroLibretto" valid:"Required"`
	NTelaio              string `form:"numeroTelaio" valid:"Required"`
	Marca                string `form:"marca" valid:"Required"`
	Modello              string `form:"modello" valid:"Required"`
	Carburante           int    `form:"carburante" valid:"Required"`
	NorEuro              int    `form:"normativaEuro" valid:"Required"`
	KW                   int    `form:"kw" valid:"Required"`
	Cilindrata           int    `form:"cilindrata" valid:"Required"`
	ConsumoTeorico       int    `form:"consumoTeorico" valid:"Required"`
	KMAnno               int    `form:"kmAnno" valid:"Required"`
	Costokm              int    `form:"costoKm" valid:"Required"`
	Pneumatici           string `form:"pneumatici" valid:"Required"`
	Contratto            int    `form:"contratto" valid:"Required"`
	NContratto           string `form:"numerocontratto" valid:"Required"`
}

//VALUTARE QUALI SONO REALMETNE NECESSARI RISPETTO AI DB
type contrattoAcq struct {
	NContratto        string  `form:"ncontratto" valid:"Required"`
	DataCont          string  `form:"datacont" valid:"Required"`
	Importo           float64 `form:"importo" valid:"Required"`
	AmmortamentoAnnuo int     `form:"ammortannuo" valid:"Required"`
	FineGaranzia      string  `form:"finegaranzia" valid:"Required"`
	KmAcquisto        int     `form:"kmacq" valid:"Required"`
	KmInizioGest      int     `form:"kmingest" valid:"Required"`
	Note              string  `form:"note" valid:"MaxSize(100)"`
	Fornitori         string  `form:"fornitore"`
	PIFornitori       string  `form:"pifornitore"`
}

type contrattoLea struct {
	NContratto   string  `form:"ncontratto" valid:"Required"`
	DataCont     string  `form:"datacont" valid:"Required"`
	PrimaRata    float64 `form:"prata" valid:"Required"`
	RataSucc     float64 `form:"ratesucc" valid:"Required"`
	NRate        int     `form:"nrate" valid:"Required"`
	Riscatto     float64 `form:"riscatto" valid:"Required"`
	DataRiscatto string  `form:"datariscatto" valid:"Required"`
	ImportoTot   float64 `form:"importotot" valid:"Required"`
	FineCont     string  `form:"datafinecontr" valid:"Required"`
	FineGaranzia string  `form:"datafinegaranzia" valid:"Required"`
	KmInizioGest int     `form:"kmingest" valid:"Required"`
	KmFineGest   int     `form:"kmfinegest" valid:"Required"`
	Note         string  `form:"note" valid:"MaxSize(100)"`
	Fornitori    string  `form:"fornitore"`
	PIFornitori  string  `form:"pifornitore"`
}

type contrattoNol struct {
	ID                  int
	NContratto          string  `form:"ncontratto" valid:"Required"`
	DataCont            string  `form:"datacont" valid:"Required"`
	DataInizio          string  `form:"datainiziocontr" valid:"Required"`
	DataFine            string  `form:"datafinecontratto" valid:"Required"`
	Riparametrizzazione int     `form:"riparametrizzazione" valid:"Required"`
	NRate               int     `form:"nrate" valid:"Required"`
	CanoneBase          float64 `form:"canonebase" valid:"Required"`
	CanoneServizi       float64 `form:"canoneservizi" valid:"Required"`
	CanoneAltro         float64 `form:"canonealtro" valid:"Required"`
	CanoneTot           float64 `form:"canonetotale" valid:"Required"`
	KmContrattuali      int     `form:"kmcontrattuali" valid:"Required"`
	AddebitoKmExtra     int     `form:"addebitokmextra" valid:"Required"`
	ImportoKm           float64 `form:"importokm" valid:"Required"`
	ImportoTot          float64 `form:"imptotale" valid:"Required"`
	KmInizioGest        int     `form:"kmingest" valid:"Required"`
	KmFineGest          int     `form:"kmfinegest" valid:"Required"`
	Note                string  `form:"note" valid:"MaxSize(100)"`
	Fornitori           string  `form:"fornitore" `
	PIFornitori         string  `form:"pifornitore" `
}

type conducente struct {
	Nome    string `form:"condnome" valid:"Required"`
	Cognome string `form:"condcognome" valid:"Required"`
	CF      string `form:"condcf" valid:"MinSize(16)"`
}

type fornitore struct {
	Descrizione string `form:"fornitore" valid:"Required"`
	PI          string `form:"pi"`
}

type data1 struct {
	dataInFlotta string `form:"dataInFlotta" valid:"Required"`
}

type dataContratto struct {
	NContratto    string `form:"numcontratto" valid:"Required"`
	TipoContratto int    `form:"tipocontratto" valid:"Required"`
}

//activeContnent function build the page
func (c *AutomezziController) activeContent(view string) {
	c.Layout = "basic-layout.tpl"
	c.Data["domainname"] = beego.AppConfig.String("appcfgdomainname")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.tpl"
	c.LayoutSections["Footer"] = "footer.tpl"
	c.TplName = view + ".tpl"

	sess := c.GetSession("portale")
	if sess != nil {
		c.Data["InSession"] = 1 // for login bar in header.tpl
		m := sess.(map[string]interface{})
		c.Data["First"] = m["first"]
		c.Data["Admin"] = m["admin"]
		c.Data["IDkey"] = m["idkey"]
		fmt.Println(m["portale"])
		c.Data["Automezzi"] = m["automezzi"]
	}
}

func (c *AutomezziController) setCompare(query string, table string) (orm.QuerySeter, bool) {

	o := orm.NewOrm()
	qs := o.QueryTable(table)
	//XSRF attack defense
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if c.Ctx.Input.Method() == "POST" {
		f := compareform{}
		if err := c.ParseForm(&f); err != nil {
			fmt.Println("cannot parse form")
			return qs, false
		}
		valid := validation.Validation{}
		if b, _ := valid.Valid(&f); !b {
			c.Data["Errors"] = valid.ErrorsMap
			return qs, false
		}
		if len(f.Compareop) >= 5 && f.Compareop[:5] == "__not" {
			qs = qs.Exclude(f.Comparefield+f.Compareop[5:], f.Compareval)
		} else {
			qs = qs.Filter(f.Comparefield+f.Compareop, f.Compareval)
		}
		c.Data["query"] = f.Comparefield + f.Compareop + "," + f.Compareval
	} else {
		str := strings.Split(query, ",")
		i := strings.Index(str[0], "__")
		if len(str[0][i:]) >= 5 && str[0][i:i+5] == "__not" {
			qs = qs.Exclude(str[0][:i]+str[0][i+5:], str[1])
		} else {
			qs = qs.Filter(str[0], str[1])
		}
		c.Data["query"] = query
	}
	return qs, true
}

//stringInSlice funcion check if string is in slice
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//ConvertInt function convert string in int
func ConvertInt(s string) int {
	//convert string in int
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

type compareform struct {
	Comparefield string `form:"comparefield"`
	Compareop    string `form:"compareop"`
	Compareval   string `form:"compareval" valid:"Required"`
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

//InitializeModule inizializza le tabelle usate dal modulo Automezzi
func InitializeModule() {

	beego.Debug("Inizializzo il modulo automezzi")

	//inizialize db automezzi
	o := orm.NewOrm()

	//CONDITION
	var maps []orm.Params
	num, err := o.QueryTable("condizioni").Values(&maps, "ID")
	if err != nil {
		fmt.Println("Errore inizializzazione tabella Condizioni")
		return
	}
	if num == 0 {
		condizioniArray := []string{"Buono stato", "Cattivo Stato", "Discreto Stato", "In Attesa di Alienazione", "In Attesa di Assegnazione", "In attesa di Riparazione", "Non utilizzabile", "Rubato", "Alienato"}

		for i := range condizioniArray {
			condizioni := models.Condizioni{ID: i + 1, Descrizione: condizioniArray[i]}

			_, err = o.Insert(&condizioni)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}
	num = 0
	err = nil
	//SPESA
	num, err = o.QueryTable("tipi_spesa").Values(&maps, "ID")
	if err != nil {
		fmt.Println("Errore inizializzazione tabella spesa")
		return
	}
	if num == 0 {

		tipiSpesaArray := []string{"Alienazione", "Assicurazione", "Bollo", "Contratto Canone", "Contratto varie", "Lavaggio", "Manutenzione ordinaria", "Pneumatici", "Revisione", "Riparazione per sinistro", "Riparazione straordinaria", "Varie"}
		for i := range tipiSpesaArray {

			tipiSpesa := models.TipiSpesa{ID: i + 1, Descrizione: tipiSpesaArray[i]}

			_, err = o.Insert(&tipiSpesa)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}
	num = 0
	err = nil

	//TIPO VEICOLO

	num, err = o.QueryTable("tipi_veicolo").Values(&maps, "ID")
	if err != nil {
		fmt.Println("Errore inizializzazione tabella vehicle_type")
		return
	}
	if num == 0 {
		tipiVeicoloArray := []string{"Autoveicolo", "Camion", "Ciclomotore", "Furgone", "Pullman", "Motoveicolo", "Altro"}

		for i := range tipiVeicoloArray {

			tipiVeicolo := models.TipiVeicolo{ID: i + 1, Descrizione: tipiVeicoloArray[i]}
			_, err = o.Insert(&tipiVeicolo)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}
	num = 0
	err = nil

	//Sector

	num, err = o.QueryTable("settori").Values(&maps, "ID")
	if err != nil {
		fmt.Println("Errore inizializzazione tabella Settori")
		return
	}
	if num == 0 {

		settoriArray := []string{"Food", "Lavanderia", "Pulizia", "Marketing", "Officina", "Agenti", "Direzione"}

		for i := range settoriArray {

			settori := models.Settori{ID: i + 1, Descrizione: settoriArray[i]}
			_, err = o.Insert(&settori)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}
	num = 0
	err = nil

	//Employment

	num, err = o.QueryTable("impieghi").Values(&maps, "ID")
	if err != nil {
		fmt.Println("Errore inizializzazione tabella Employment")
		return
	}
	if num == 0 {

		impiegoArray := []string{"Aziendale", "Aziendale + Personale", "Personale"}

		for i := range impiegoArray {

			impiego := models.Impieghi{ID: i + 1, Descrizione: impiegoArray[i]}
			_, err = o.Insert(&impiego)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}
	num = 0
	err = nil

	//Responsabilita

	num, err = o.QueryTable("responsabilita").Values(&maps, "ID")
	if err != nil {
		fmt.Println("Errore inizializzazione tabella Responsabilita")
		return
	}
	if num == 0 {

		responsabilitaArray := []string{"Concorso di colpa", "Da accertare", "Della controparte", "Propria"}

		for i := range responsabilitaArray {

			responsabilita := models.Responsabilita{ID: i + 1, Descrizione: responsabilitaArray[i]}
			_, err = o.Insert(&responsabilita)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}
	num = 0
	err = nil

	//tipo_infrazione

	num, err = o.QueryTable("tipi_infrazione").Values(&maps, "ID")
	if err != nil {
		fmt.Println("Errore inizializzazione tabella tipo_infrazione")
		return
	}
	if num == 0 {

		tipiInfrazioneArray := []string{"Accesso in senso vietato", "Accesso in senso vietato", "Cinture di sicurezza non allacciate", "Eccesso di velocità", "Guida contromano", "Guida pericolosa", "Positivo a controlllo alcool", "Precedenza non rispettata", "Semaforo rosso", "Sosta vietata", "Utilizzo di telefono cellulare", "Violazione di corsia preferenziale", "Violazione di ztl", "Altro"}

		for i := range tipiInfrazioneArray {

			tipiInfrazione := models.TipiInfrazione{ID: i + 1, Descrizione: tipiInfrazioneArray[i]}
			_, err = o.Insert(&tipiInfrazione)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}
	num = 0
	err = nil

	//Carburante

	num, err = o.QueryTable("carburante").Values(&maps, "ID")
	if err != nil {
		fmt.Println("Errore inizializzazione tabella Carburante")
		return
	}
	if num == 0 {

		carburanteArray := []string{"Benzina", "Diesel", "Gas", "Metano"}

		for i := range carburanteArray {

			carburante := models.Carburante{ID: i + 1, Descrizione: carburanteArray[i]}
			_, err = o.Insert(&carburante)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}
	num = 0
	err = nil

}

//AddFornitore permette di aggiungere un nuovo fornitore
func (c *AutomezziController) AddFornitore() {
	c.activeContent("automezzi/add/fornitore")
	//verifica se l'utente è loggato
	sess := c.GetSession("portale")
	if sess == nil {
		c.Redirect("/", 302)
		return
	}
	//inizializza i messaggi
	flash := beego.NewFlash()
	//Estrae la sessione
	m := sess.(map[string]interface{})
	fmt.Println(m["admin"])
	fmt.Println(reflect.ValueOf(m["admin"]).Type())
	//verifica se l'utente ha i privilegi di amministratore
	if m["admin"] != 3 {
		flash.Notice("Non hai i diritti per accedere a questa pagina")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}
	fmt.Printf("hai i diritti")
	/* NON HA PARAMETRI
	parms := c.Ctx.Input.Param(":parms")
	fmt.Println(parms)
	c.Data["parms"] = parms
	*/
	//XSRF attack defense
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if c.Ctx.Input.Method() == "POST" {

		u := fornitore{}
		//legge il form
		if err := c.ParseForm(&u); err != nil {
			fmt.Println("cannot parse form")
			return
		}
		fmt.Println(u)
		c.Data["Fornitore"] = u
		//valida i dati
		valid := validation.Validation{}
		if b, _ := valid.Valid(&u); !b {
			c.Data["Errors"] = valid.ErrorsMap
			return
		}
		//******** Save user info to database
		o := orm.NewOrm()
		o.Using("default")

		//crea il fornitore

		Fornitore := models.Fornitori{Descrizione: u.Descrizione, PI: u.PI}

		//inserisce il fornitore nel db
		_, err := o.Insert(&Fornitore)
		if err != nil {
			flash.Error("Errore autorizzazioni applicazioni")
			flash.Store(&c.Controller)
			return
		}

		flash.Notice("Fornitore Registrato")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}
}

//ViewFornitori permette di visualizzare i fornitori inseriti sul portale
func (c *AutomezziController) ViewFornitori() {
	// Only administrator can Manage accounts
	c.activeContent("automezzi/view/fornitori")

	//******** c page requires login
	sess := c.GetSession("portale")
	if sess == nil {
		c.Redirect("/", 302)
		return
	}
	flash := beego.NewFlash()
	m := sess.(map[string]interface{})
	fmt.Println(m["admin"])
	fmt.Println(reflect.ValueOf(m["admin"]).Type())
	if m["admin"] != 3 {
		flash.Notice("Non hai i diritti per accedere a questa pagina")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}

	fmt.Printf("hai i diritti")

	//in caso di panic reindirizza alla home
	defer func(c *AutomezziController) {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Index", r)
			c.Redirect("/", 302)
		}
	}(c)
	//Se non ci sono parametri forzala ricerca standard
	var parms string
	if c.Ctx.Input.Param(":parms") == "" {
		parms = "id!0!id__gte,0"
	} else {
		parms = c.Ctx.Input.Param(":parms")
	}
	//parametro serve per suddividere in pagine l'elenco utenti
	const pagesize = 10
	c.Data["parms"] = parms
	str := strings.Split(parms, "!")
	fmt.Println("parms is", str)
	order := str[0]
	off, _ := strconv.Atoi(str[1])
	offset := int64(off)
	if offset < 0 {
		offset = 0
	}
	query := str[2]

	var fornitori []*models.Fornitori
	rows := ""

	qs, ok := c.setCompare(query, "fornitori")
	if !ok {
		fmt.Println("cannot set QuerySeter")
		o := orm.NewOrm()
		qs := o.QueryTable("fornitori")
		qs = qs.Filter("id__gte", 0)
		c.Data["query"] = "id__gte,0"
	}

	count, _ := qs.Count()
	c.Data["count"] = count
	if offset >= count {
		offset = 0
	}
	num, err := qs.Limit(pagesize, offset).OrderBy(order).All(&fornitori)
	if err != nil {
		fmt.Println("Query table failed:", err)
	}
	/*TABELLA IN BASE AI PARAMETRI*/
	for i := range fornitori {
		rows += fmt.Sprintf("<tr><td>%s</td>"+
			"<td>%s</td><td><a class='btn btn-defalut btn-xs' href='http://%s/automezzi/fornitore/%d'><span class='glyphicon glyphicon-edit'></span> Modifica</a></td><td><a class='btn btn-defalut btn-xs' href='http://%s/automezzi/add/acquisto/%d'><span class='glyphicon glyphicon-edit'></span> Acquisto</a> <a class='btn btn-defalut btn-xs' href='http://%s/automezzi/add/leasing/%d'><span class='glyphicon glyphicon-edit'></span> Leasing</a> <a class='btn btn-defalut btn-xs' href='http://%s/automezzi/add/noleggio/%d'><span class='glyphicon glyphicon-edit'></span> Noleggio</a></td></tr>", fornitori[i].Descrizione, fornitori[i].PI, appcfgdomainname, fornitori[i].ID, appcfgdomainname, fornitori[i].ID, appcfgdomainname, fornitori[i].ID, appcfgdomainname, fornitori[i].ID)
	}
	c.Data["Rows"] = template.HTML(rows)

	c.Data["order"] = order
	c.Data["offset"] = offset
	c.Data["end"] = max(0, (count/pagesize)*pagesize)

	if num+offset < count {
		c.Data["next"] = num + offset
	}
	if offset-pagesize >= 0 {
		c.Data["prev"] = offset - pagesize
		c.Data["showprev"] = true
	} else if offset > 0 && offset < pagesize {
		c.Data["prev"] = 0
		c.Data["showprev"] = true
	}

	if count > pagesize {
		c.Data["ShowNav"] = true
	}
	c.Data["progress"] = float64(offset*100) / float64(max(count, 1))

}

//ViewContrattiAcq permette di visualizzare i contratti di acquisto inseriti sul portale
func (c *AutomezziController) ViewContrattiAcq() {
	// Only administrator can Manage accounts
	c.activeContent("automezzi/view/contrattiacq")

	//******** c page requires login
	sess := c.GetSession("portale")
	if sess == nil {
		c.Redirect("/", 302)
		return
	}
	flash := beego.NewFlash()
	m := sess.(map[string]interface{})
	fmt.Println(m["admin"])
	fmt.Println(reflect.ValueOf(m["admin"]).Type())
	if m["admin"] != 3 {
		flash.Notice("Non hai i diritti per accedere a questa pagina")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}

	fmt.Printf("hai i diritti")

	//in caso di panic reindirizza alla home
	defer func(c *AutomezziController) {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Index", r)
			c.Redirect("/", 302)
		}
	}(c)
	//NON VA SENZA PARAMETRI
	//parametro serve per suddividere in pagine l'elenco contratti e se non è presente viene forzato
	const pagesize = 10
	var parms string
	if c.Ctx.Input.Param(":parms") == "" {
		parms = "id!0!id__gte,0"
	} else {
		parms = c.Ctx.Input.Param(":parms")
	}

	c.Data["parms"] = parms
	str := strings.Split(parms, "!")
	fmt.Println("parms is", str)
	order := str[0]
	off, _ := strconv.Atoi(str[1])
	offset := int64(off)
	if offset < 0 {
		offset = 0
	}
	query := str[2]

	rows := ""
	var contratti []*models.ContrAcquisti
	qs, ok := c.setCompare(query, "contr_acquisti")
	if !ok {
		fmt.Println("cannot set QuerySeter")
		o := orm.NewOrm()
		qs := o.QueryTable("contr_acquisti")
		qs = qs.Filter("id__gte", 0)
		c.Data["query"] = "id__gte,0"
	}

	count, _ := qs.Count()
	c.Data["count"] = count
	if offset >= count {
		offset = 0
	}

	num, err := qs.Limit(pagesize, offset).OrderBy(order).RelatedSel().All(&contratti)
	if err != nil {
		fmt.Println("Query table failed:", err)
	}

	/*TABELLA IN BASE AI PARAMETRI*/
	for i := range contratti {

		rows += fmt.Sprintf("<tr><td>%s</td>"+
			"<td>%s</td><td>%s</td><td><a class='btn btn-defalut btn-xs' href='http://%s/automezzi/edit/contrattoacq/%d'><span class='glyphicon glyphicon-edit'></span> Modifica</a></td></tr>", contratti[i].NContratto, contratti[i].DataCont, contratti[i].Fornitori.Descrizione, appcfgdomainname, contratti[i].ID)
	}
	c.Data["Rows"] = template.HTML(rows)

	c.Data["order"] = order
	c.Data["offset"] = offset
	c.Data["end"] = max(0, (count/pagesize)*pagesize)

	if num+offset < count {
		c.Data["next"] = num + offset
	}
	if offset-pagesize >= 0 {
		c.Data["prev"] = offset - pagesize
		c.Data["showprev"] = true
	} else if offset > 0 && offset < pagesize {
		c.Data["prev"] = 0
		c.Data["showprev"] = true
	}

	if count > pagesize {
		c.Data["ShowNav"] = true
	}
	c.Data["progress"] = float64(offset*100) / float64(max(count, 1))

}

//EditAcquisto permette di modificare un nuovo contratto di acquisto
//Il parametro passato alla pagina è ContrAcquisti.ID
func (c *AutomezziController) EditAcquisto() {
	c.activeContent("automezzi/edit/contrattoacq")
	//verifica se l'utente è loggato
	sess := c.GetSession("portale")
	if sess == nil {
		c.Redirect("/", 302)
		return
	}
	//inizializza i messaggi
	flash := beego.NewFlash()
	//carica la sessione
	m := sess.(map[string]interface{})
	//verifica se si detengono i privilegi per accedere alla pagina
	if m["admin"] != 3 {
		flash.Notice("Non hai i diritti per accedere a questa pagina")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}
	fmt.Printf("hai i diritti")
	//prende i paramentri della pagina nel formato Ncontratto!IDFornitore
	parms := c.Ctx.Input.Param(":parms")
	idcontr, _ := strconv.Atoi(parms)
	c.Data["ID"] = idcontr
	contratto := models.ContrAcquisti{ID: idcontr}

	o := orm.NewOrm()
	o.Using("default")
	//Flag per stabilire se il contratto è nuovo o esisteva già

	err := o.QueryTable("contr_acquisti").Filter("ID", idcontr).RelatedSel().One(&contratto)
	if err == orm.ErrNoRows {
		//non esiste nessun contratto con questo codice
		flash.Notice("In contratto di acquisto non esiste")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	} else if err == orm.ErrMissPK {
		fmt.Println("Errore - Contattare l'amministratore del sito")
	}
	//aggiorna il form con i dati letti dal DB

	defer func(c *AutomezziController, contratto *models.ContrAcquisti) {

		c.Data["ID"] = contratto.ID
		c.Data["NContratto"] = contratto.NContratto
		c.Data["DataCont"] = contratto.DataCont.Format("2006-01-02")
		c.Data["Importo"] = contratto.Importo
		c.Data["AmmortamentoAnnuo"] = contratto.AmmortamentoAnnuo
		c.Data["FineGaranzia"] = contratto.FineGaranzia.Format("2006-01-02")
		c.Data["KmAcquisto"] = contratto.KmAcquisto
		c.Data["KmInizioGest"] = contratto.KmInizioGest
		c.Data["Note"] = contratto.Note
		c.Data["Fornitori"] = contratto.Fornitori.Descrizione
		c.Data["PIFornitori"] = contratto.Fornitori.PI

	}(c, &contratto)

	//XSRF attack defense
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if c.Ctx.Input.Method() == "POST" {

		u := contrattoAcq{}

		if err := c.ParseForm(&u); err != nil {
			fmt.Println("cannot parse form")
			return
		}
		f := models.Fornitori{ID: contratto.Fornitori.ID, Descrizione: u.Fornitori, PI: u.PIFornitori}
		fmt.Println(u.NContratto)
		fmt.Println(u.DataCont)
		fmt.Println(f.Descrizione)
		c.Data["data"] = u
		valid := validation.Validation{}
		if b, _ := valid.Valid(&u); !b {
			c.Data["Errors"] = valid.ErrorsMap
			return
		}
		//Aggiorno Fornitore
		_, err := o.Update(&f)
		if err != nil {
			flash.Error("Errore interno")
			flash.Store(&c.Controller)
			return
		}
		//Converte le data da string a time
		dcont, err := time.Parse("2006-01-02 15:04", u.DataCont+" 10:00")
		if err != nil {
			fmt.Println(err)
			return
		}
		dgar, err := time.Parse("2006-01-02 15:04", u.DataCont+" 10:00")
		if err != nil {
			fmt.Println(err)
			return
		}

		contratto.NContratto = u.NContratto
		contratto.DataCont = dcont
		contratto.Importo = u.Importo
		contratto.AmmortamentoAnnuo = u.AmmortamentoAnnuo
		contratto.FineGaranzia = dgar
		contratto.KmAcquisto = u.KmAcquisto
		contratto.KmInizioGest = u.KmInizioGest
		contratto.Note = u.Note

		//Aggiorno Contratto
		_, err = o.Update(&contratto)
		if err != nil {
			flash.Error("Errore interno nell'aggiornare il contratto")
			flash.Store(&c.Controller)
			return
		}

		flash.Notice("Data Registrato")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)

	}
}

//AddAcquisto permette di caricare un nuovo contratto di acquisto
//Il parametro passato alla pagina è Fornitore.ID
func (c *AutomezziController) AddAcquisto() {
	c.activeContent("automezzi/add/acquisto")
	//verifica se l'utente è loggato
	sess := c.GetSession("portale")
	if sess == nil {
		c.Redirect("/", 302)
		return
	}
	//inizializza i messaggi
	flash := beego.NewFlash()
	//carica la sessione
	m := sess.(map[string]interface{})
	//verifica se si detengono i privilegi per accedere alla pagina
	if m["admin"] != 3 {
		flash.Notice("Non hai i diritti per accedere a questa pagina")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}
	fmt.Printf("hai i diritti")
	//prende i paramentri della pagina nel formato Ncontratto!IDFornitore
	parms := c.Ctx.Input.Param(":parms")
	idforn, err := strconv.Atoi(parms)
	if err != nil {
		fmt.Println(err)
		flash.Error("Impossibile convertire in Int il codice fornitore")
		flash.Store(&c.Controller)
		return
	}
	//Inizializza il contratto
	contratto := models.ContrAcquisti{}

	//Estrae i dati del fornitore partendo dal ID passato come parametro
	fornitore := models.Fornitori{ID: idforn}

	o := orm.NewOrm()
	o.Using("default")
	err = o.Read(&fornitore, "ID")
	if err != nil {
		fmt.Println(err)
		flash.Error("Internal error")
		flash.Store(&c.Controller)
		return
	}
	//Inserisce nel form i dati del fornitore
	defer func(c *AutomezziController, fornitore *models.Fornitori) {

		c.Data["Fornitori"] = fornitore.Descrizione
		c.Data["PIFornitori"] = fornitore.PI

	}(c, &fornitore)

	//XSRF attack defense
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if c.Ctx.Input.Method() == "POST" {

		u := contrattoAcq{}
		//Estra i dati dal form
		if err := c.ParseForm(&u); err != nil {
			fmt.Println("cannot parse form")
			return
		}

		//c.Data["data"] = u
		//Valida i dati del form
		valid := validation.Validation{}
		if b, _ := valid.Valid(&u); !b {
			c.Data["Errors"] = valid.ErrorsMap
			return
		}

		//Trasforma le date da string a Time
		dcont, err := time.Parse("2006-01-02 15:04", u.DataCont+" 10:00")
		if err != nil {
			fmt.Println(err)
			return
		}
		dgar, err := time.Parse("2006-01-02 15:04", u.FineGaranzia+" 10:00")
		if err != nil {
			fmt.Println(err)
			return
		}
		//Inserisce i dati estratti dal form den contratto
		contratto.NContratto = u.NContratto
		contratto.DataCont = dcont
		contratto.Importo = u.Importo
		contratto.AmmortamentoAnnuo = u.AmmortamentoAnnuo
		contratto.FineGaranzia = dgar
		contratto.KmAcquisto = u.KmAcquisto
		contratto.KmInizioGest = u.KmInizioGest
		contratto.Note = u.Note
		contratto.Fornitori = &fornitore

		//inserisce il contratto nel db
		_, err = o.Insert(&contratto)
		if err != nil {
			flash.Error("Errore autorizzazioni applicazioni")
			flash.Store(&c.Controller)
			return
		}

		flash.Notice("Contratto Registrato")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)

	}
}

//ViewContrattiLea permette di visualizzare i contratti di leasing inseriti sul portale
func (c *AutomezziController) ViewContrattiLea() {
	// Only administrator can Manage accounts
	c.activeContent("automezzi/view/contrattilea")

	//******** c page requires login
	sess := c.GetSession("portale")
	if sess == nil {
		c.Redirect("/", 302)
		return
	}
	flash := beego.NewFlash()
	m := sess.(map[string]interface{})
	fmt.Println(m["admin"])
	fmt.Println(reflect.ValueOf(m["admin"]).Type())
	if m["admin"] != 3 {
		flash.Notice("Non hai i diritti per accedere a questa pagina")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}

	fmt.Printf("hai i diritti")

	//in caso di panic reindirizza alla home
	defer func(c *AutomezziController) {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Index", r)
			c.Redirect("/", 302)
		}
	}(c)
	//NON VA SENZA PARAMETRI
	//parametro serve per suddividere in pagine l'elenco contratti e se non è presente viene forzato
	const pagesize = 10
	var parms string
	if c.Ctx.Input.Param(":parms") == "" {
		parms = "id!0!id__gte,0"
	} else {
		parms = c.Ctx.Input.Param(":parms")
	}

	c.Data["parms"] = parms
	str := strings.Split(parms, "!")
	fmt.Println("parms is", str)
	order := str[0]
	off, _ := strconv.Atoi(str[1])
	offset := int64(off)
	if offset < 0 {
		offset = 0
	}
	query := str[2]

	rows := ""
	var contratti []*models.ContrLeasing
	qs, ok := c.setCompare(query, "contr_leasing")
	if !ok {
		fmt.Println("cannot set QuerySeter")
		o := orm.NewOrm()
		qs := o.QueryTable("contr_leasing")
		qs = qs.Filter("id__gte", 0)
		c.Data["query"] = "id__gte,0"
	}

	count, _ := qs.Count()
	c.Data["count"] = count
	if offset >= count {
		offset = 0
	}

	num, err := qs.Limit(pagesize, offset).OrderBy(order).RelatedSel().All(&contratti)
	if err != nil {
		fmt.Println("Query table failed:", err)
	}

	/*TABELLA IN BASE AI PARAMETRI*/
	for i := range contratti {
		rows += fmt.Sprintf("<tr><td>%s</td>"+
			"<td>%s</td><td>%s</td><td><a class='btn btn-defalut btn-xs' href='http://%s/automezzi/edit/contrattolea/%d'><span class='glyphicon glyphicon-edit'></span> Modifica</a></td></tr>", contratti[i].NContratto, contratti[i].DataCont, contratti[i].Fornitori.Descrizione, appcfgdomainname, contratti[i].ID)
	}
	c.Data["Rows"] = template.HTML(rows)

	c.Data["order"] = order
	c.Data["offset"] = offset
	c.Data["end"] = max(0, (count/pagesize)*pagesize)

	if num+offset < count {
		c.Data["next"] = num + offset
	}
	if offset-pagesize >= 0 {
		c.Data["prev"] = offset - pagesize
		c.Data["showprev"] = true
	} else if offset > 0 && offset < pagesize {
		c.Data["prev"] = 0
		c.Data["showprev"] = true
	}

	if count > pagesize {
		c.Data["ShowNav"] = true
	}
	c.Data["progress"] = float64(offset*100) / float64(max(count, 1))

}

//EditLeasing permette di modificare un nuovo contratto di leasing
//Il parametro passato alla pagina è ContrLeasing.ID
func (c *AutomezziController) EditLeasing() {
	c.activeContent("automezzi/edit/contrattoacq")
	//verifica se l'utente è loggato
	sess := c.GetSession("portale")
	if sess == nil {
		c.Redirect("/", 302)
		return
	}
	//inizializza i messaggi
	flash := beego.NewFlash()
	//carica la sessione
	m := sess.(map[string]interface{})
	//verifica se si detengono i privilegi per accedere alla pagina
	if m["admin"] != 3 {
		flash.Notice("Non hai i diritti per accedere a questa pagina")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}
	fmt.Printf("hai i diritti")
	//prende i paramentri della pagina nel formato Ncontratto!IDFornitore
	parms := c.Ctx.Input.Param(":parms")
	idcontr, _ := strconv.Atoi(parms)
	c.Data["ID"] = idcontr
	contratto := models.ContrLeasing{ID: idcontr}

	o := orm.NewOrm()
	o.Using("default")
	//Flag per stabilire se il contratto è nuovo o esisteva già

	err := o.QueryTable("contr_leasing").Filter("ID", idcontr).RelatedSel().One(&contratto)
	if err == orm.ErrNoRows {
		//non esiste nessun contratto con questo codice
		flash.Notice("In contratto di contr_leasing non esiste")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	} else if err == orm.ErrMissPK {
		fmt.Println("Errore - Contattare l'amministratore del sito")
	}
	//aggiorna il form con i dati letti dal DB

	defer func(c *AutomezziController, contratto *models.ContrLeasing) {

		c.Data["ID"] = contratto.ID
		c.Data["NContratto"] = contratto.NContratto
		c.Data["DataCont"] = contratto.DataCont.Format("2006-01-02")
		c.Data["PrimaRata"] = contratto.PrimaRata
		c.Data["RataSucc"] = contratto.RataSucc
		c.Data["NRate"] = contratto.NRate
		c.Data["Riscatto"] = contratto.Riscatto
		c.Data["DataRiscatto"] = contratto.DataRiscatto.Format("2006-01-02")
		c.Data["ImportoTot"] = contratto.ImportoTot
		c.Data["FineCont"] = contratto.FineCont.Format("2006-01-02")
		c.Data["FineGaranzia"] = contratto.FineGaranzia.Format("2006-01-02")
		c.Data["KmInizioGest"] = contratto.KmInizioGest
		c.Data["KmFineGest"] = contratto.KmFineGest
		c.Data["Note"] = contratto.Note
		c.Data["Fornitori"] = contratto.Fornitori.Descrizione
		c.Data["PIFornitori"] = contratto.Fornitori.PI

	}(c, &contratto)

	//XSRF attack defense
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if c.Ctx.Input.Method() == "POST" {

		u := contrattoLea{}

		if err := c.ParseForm(&u); err != nil {
			fmt.Println("cannot parse form")
			return
		}
		f := models.Fornitori{ID: contratto.Fornitori.ID, Descrizione: u.Fornitori, PI: u.PIFornitori}
		fmt.Println(u.NContratto)
		fmt.Println(u.DataCont)
		fmt.Println(f.Descrizione)
		c.Data["data"] = u
		valid := validation.Validation{}
		if b, _ := valid.Valid(&u); !b {
			c.Data["Errors"] = valid.ErrorsMap
			return
		}
		//Aggiorno Fornitore
		_, err := o.Update(&f)
		if err != nil {
			flash.Error("Errore interno")
			flash.Store(&c.Controller)
			return
		}
		//Converte le data da string a time
		dcont, err := time.Parse("2006-01-02 15:04", u.DataCont+" 10:00")
		if err != nil {
			fmt.Println(err)
			return
		}
		dris, err := time.Parse("2006-01-02 15:04", u.DataRiscatto+" 10:00")
		if err != nil {
			fmt.Println(err)
			return
		}
		dfcon, err := time.Parse("2006-01-02 15:04", u.FineCont+" 10:00")
		if err != nil {
			fmt.Println(err)
			return
		}
		dgar, err := time.Parse("2006-01-02 15:04", u.FineGaranzia+" 10:00")
		if err != nil {
			fmt.Println(err)
			return
		}

		contratto.NContratto = u.NContratto
		contratto.DataCont = dcont
		contratto.PrimaRata = u.PrimaRata
		contratto.RataSucc = u.RataSucc
		contratto.NRate = u.NRate
		contratto.Riscatto = u.Riscatto
		contratto.DataRiscatto = dris
		contratto.ImportoTot = u.ImportoTot
		contratto.FineCont = dfcon
		contratto.FineGaranzia = dgar
		contratto.KmInizioGest = u.KmInizioGest
		contratto.KmFineGest = u.KmFineGest
		contratto.Note = u.Note

		//Aggiorno Contratto
		_, err = o.Update(&contratto)
		if err != nil {
			flash.Error("Errore interno nell'aggiornare il contratto")
			flash.Store(&c.Controller)
			return
		}

		flash.Notice("Contratto Leasing Modificato")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)

	}
}

//AddLeasing permette di caricare un nuovo contratto di leasing
//Il parametro passato alla pagina è Fornitore.ID
func (c *AutomezziController) AddLeasing() {
	c.activeContent("automezzi/add/leasing")
	//verifica se l'utente è loggato
	sess := c.GetSession("portale")
	if sess == nil {
		c.Redirect("/", 302)
		return
	}
	//inizializza i messaggi
	flash := beego.NewFlash()
	//carica la sessione
	m := sess.(map[string]interface{})
	//verifica se si detengono i privilegi per accedere alla pagina
	if m["admin"] != 3 {
		flash.Notice("Non hai i diritti per accedere a questa pagina")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}
	fmt.Printf("hai i diritti")
	//prende i paramentri della pagina nel formato Ncontratto!IDFornitore
	parms := c.Ctx.Input.Param(":parms")
	idforn, err := strconv.Atoi(parms)
	if err != nil {
		fmt.Println(err)
		flash.Error("Impossibile convertire in Int il codice fornitore")
		flash.Store(&c.Controller)
		return
	}
	//Inizializza il contratto
	contratto := models.ContrLeasing{}

	//Estrae i dati del fornitore partendo dal ID passato come parametro
	fornitore := models.Fornitori{ID: idforn}

	o := orm.NewOrm()
	o.Using("default")
	err = o.Read(&fornitore, "ID")
	if err != nil {
		fmt.Println(err)
		flash.Error("Internal error")
		flash.Store(&c.Controller)
		return
	}
	//Inserisce nel form i dati del fornitore
	defer func(c *AutomezziController, fornitore *models.Fornitori) {

		c.Data["Fornitori"] = fornitore.Descrizione
		c.Data["PIFornitori"] = fornitore.PI

	}(c, &fornitore)

	//XSRF attack defense
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if c.Ctx.Input.Method() == "POST" {

		u := contrattoLea{}
		//Estra i dati dal form
		if err := c.ParseForm(&u); err != nil {
			fmt.Println("cannot parse form")
			return
		}

		//c.Data["data"] = u
		//Valida i dati del form
		valid := validation.Validation{}
		if b, _ := valid.Valid(&u); !b {
			c.Data["Errors"] = valid.ErrorsMap
			return
		}

		//Trasforma le date da string a Time
		fmt.Println(u.DataCont)
		dcont, err := time.Parse("2006-01-02 15:04", u.DataCont+" 10:00")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(u.DataRiscatto)
		dris, err := time.Parse("2006-01-02 15:04", u.DataRiscatto+" 10:00")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(u.FineCont)
		dfcon, err := time.Parse("2006-01-02 15:04", u.FineCont+" 10:00")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(u.FineGaranzia)
		dgar, err := time.Parse("2006-01-02 15:04", u.FineGaranzia+" 10:00")
		if err != nil {
			fmt.Println(err)
			return
		}
		//Inserisce i dati estratti dal form den contratto

		contratto.NContratto = u.NContratto
		contratto.DataCont = dcont
		contratto.PrimaRata = u.PrimaRata
		contratto.RataSucc = u.RataSucc
		contratto.NRate = u.NRate
		contratto.Riscatto = u.Riscatto
		contratto.DataRiscatto = dris
		contratto.ImportoTot = u.ImportoTot
		contratto.FineCont = dfcon
		contratto.FineGaranzia = dgar
		contratto.KmInizioGest = u.KmInizioGest
		contratto.KmFineGest = u.KmFineGest
		contratto.Note = u.Note
		contratto.Fornitori = &fornitore

		//inserisce il contratto nel db
		_, err = o.Insert(&contratto)
		if err != nil {
			flash.Error("Errore autorizzazioni applicazioni")
			flash.Store(&c.Controller)
			return
		}

		flash.Notice("Contratto Leasing Registrato")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)

	}
}

//ViewContrattiNol permette di visualizzare i contratti di noleggio inseriti sul portale
func (c *AutomezziController) ViewContrattiNol() {
	// Only administrator can Manage accounts
	c.activeContent("automezzi/view/contrattinol")

	//******** c page requires login
	sess := c.GetSession("portale")
	if sess == nil {
		c.Redirect("/", 302)
		return
	}
	flash := beego.NewFlash()
	m := sess.(map[string]interface{})
	fmt.Println(m["admin"])
	fmt.Println(reflect.ValueOf(m["admin"]).Type())
	if m["admin"] != 3 {
		flash.Notice("Non hai i diritti per accedere a questa pagina")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}

	fmt.Printf("hai i diritti")

	//in caso di panic reindirizza alla home
	defer func(c *AutomezziController) {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Index", r)
			c.Redirect("/", 302)
		}
	}(c)
	//NON VA SENZA PARAMETRI
	//parametro serve per suddividere in pagine l'elenco contratti e se non è presente viene forzato
	const pagesize = 10
	var parms string
	if c.Ctx.Input.Param(":parms") == "" {
		parms = "id!0!id__gte,0"
	} else {
		parms = c.Ctx.Input.Param(":parms")
	}

	c.Data["parms"] = parms
	str := strings.Split(parms, "!")
	fmt.Println("parms is", str)
	order := str[0]
	off, _ := strconv.Atoi(str[1])
	offset := int64(off)
	if offset < 0 {
		offset = 0
	}
	query := str[2]

	rows := ""
	var contratti []*models.ContrNoleggi
	qs, ok := c.setCompare(query, "contr_noleggi")
	if !ok {
		fmt.Println("cannot set QuerySeter")
		o := orm.NewOrm()
		qs := o.QueryTable("contr_noleggi")
		qs = qs.Filter("id__gte", 0)
		c.Data["query"] = "id__gte,0"
	}

	count, _ := qs.Count()
	c.Data["count"] = count
	if offset >= count {
		offset = 0
	}

	num, err := qs.Limit(pagesize, offset).OrderBy(order).RelatedSel().All(&contratti)
	if err != nil {
		fmt.Println("Query table failed:", err)
	}

	/*TABELLA IN BASE AI PARAMETRI*/
	for i := range contratti {

		rows += fmt.Sprintf("<tr><td>%s</td>"+
			"<td>%s</td><td>%s</td><td><a class='btn btn-defalut btn-xs' href='http://%s/automezzi/edit/contrattonol/%d'><span class='glyphicon glyphicon-edit'></span> Modifica</a></td></tr>", contratti[i].NContratto, contratti[i].DataCont, contratti[i].Fornitori.Descrizione, appcfgdomainname, contratti[i].ID)
	}
	c.Data["Rows"] = template.HTML(rows)

	c.Data["order"] = order
	c.Data["offset"] = offset
	c.Data["end"] = max(0, (count/pagesize)*pagesize)

	if num+offset < count {
		c.Data["next"] = num + offset
	}
	if offset-pagesize >= 0 {
		c.Data["prev"] = offset - pagesize
		c.Data["showprev"] = true
	} else if offset > 0 && offset < pagesize {
		c.Data["prev"] = 0
		c.Data["showprev"] = true
	}

	if count > pagesize {
		c.Data["ShowNav"] = true
	}
	c.Data["progress"] = float64(offset*100) / float64(max(count, 1))

}

//EditNoleggio permette di modificare un nuovo contratto di noleggio
//Il parametro passato alla pagina è ContrNoleggi.ID
func (c *AutomezziController) EditNoleggio() {
	c.activeContent("automezzi/edit/contrattonol")
	//verifica se l'utente è loggato
	sess := c.GetSession("portale")
	if sess == nil {
		c.Redirect("/", 302)
		return
	}
	//inizializza i messaggi
	flash := beego.NewFlash()
	//carica la sessione
	m := sess.(map[string]interface{})
	//verifica se si detengono i privilegi per accedere alla pagina
	if m["admin"] != 3 {
		flash.Notice("Non hai i diritti per accedere a questa pagina")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}
	fmt.Printf("hai i diritti")
	//prende i paramentri della pagina nel formato Ncontratto!IDFornitore
	parms := c.Ctx.Input.Param(":parms")
	idcontr, _ := strconv.Atoi(parms)
	c.Data["ID"] = idcontr
	contratto := models.ContrNoleggi{ID: idcontr}

	o := orm.NewOrm()
	o.Using("default")
	//Flag per stabilire se il contratto è nuovo o esisteva già

	err := o.QueryTable("contr_noleggi").Filter("ID", idcontr).RelatedSel().One(&contratto)
	if err == orm.ErrNoRows {
		//non esiste nessun contratto con questo codice
		flash.Notice("In contratto di Noleggio non esiste")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	} else if err == orm.ErrMissPK {
		fmt.Println("Errore - Contattare l'amministratore del sito")
	}
	//aggiorna il form con i dati letti dal DB

	defer func(c *AutomezziController, contratto *models.ContrNoleggi) {

		c.Data["ID"] = contratto.ID
		c.Data["NContratto"] = contratto.NContratto
		c.Data["DataCont"] = contratto.DataCont.Format("2006-01-02")
		c.Data["DataInizio"] = contratto.DataInizio.Format("2006-01-02")
		c.Data["DataFine"] = contratto.DataFine.Format("2006-01-02")
		c.Data["Riparametrizzazione"] = contratto.Riparametrizzazione
		c.Data["NRate"] = contratto.NRate
		c.Data["CanoneBase"] = contratto.CanoneBase
		c.Data["CanoneServizi"] = contratto.CanoneServizi
		c.Data["CanoneAltro"] = contratto.CanoneAltro
		c.Data["CanoneTot"] = contratto.CanoneTot
		c.Data["KmContrattuali"] = contratto.KmContrattuali
		c.Data["AddebitoKmExtra"] = contratto.AddebitoKmExtra
		c.Data["ImportoKm"] = contratto.ImportoKm
		c.Data["ImportoTot"] = contratto.ImportoTot
		c.Data["KmInizioGest"] = contratto.KmInizioGest
		c.Data["KmFineGest"] = contratto.KmFineGest
		c.Data["Note"] = contratto.Note
		c.Data["Fornitori"] = contratto.Fornitori.Descrizione
		c.Data["PIFornitori"] = contratto.Fornitori.PI

	}(c, &contratto)

	//XSRF attack defense
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if c.Ctx.Input.Method() == "POST" {

		u := contrattoNol{}

		if err := c.ParseForm(&u); err != nil {
			fmt.Println("cannot parse form")
			return
		}
		f := models.Fornitori{ID: contratto.Fornitori.ID, Descrizione: u.Fornitori, PI: u.PIFornitori}
		fmt.Println(u.NContratto)
		fmt.Println(u.DataCont)
		fmt.Println(f.Descrizione)
		c.Data["data"] = u
		valid := validation.Validation{}
		if b, _ := valid.Valid(&u); !b {
			c.Data["Errors"] = valid.ErrorsMap
			return
		}
		//Aggiorno Fornitore
		_, err := o.Update(&f)
		if err != nil {
			flash.Error("Errore interno")
			flash.Store(&c.Controller)
			return
		}
		//Converte le data da string a time
		dcont, err := time.Parse("2006-01-02 15:04", u.DataCont+" 10:00")
		if err != nil {
			fmt.Println(err)
			return
		}
		diniz, err := time.Parse("2006-01-02 15:04", u.DataInizio+" 10:00")
		if err != nil {
			fmt.Println(err)
			return
		}
		dfine, err := time.Parse("2006-01-02 15:04", u.DataFine+" 10:00")
		if err != nil {
			fmt.Println(err)
			return
		}

		contratto.NContratto = u.NContratto
		contratto.DataCont = dcont
		contratto.DataInizio = diniz
		contratto.DataFine = dfine
		contratto.Riparametrizzazione = u.Riparametrizzazione
		contratto.NRate = u.NRate
		contratto.CanoneBase = u.CanoneBase
		contratto.CanoneServizi = u.CanoneServizi
		contratto.CanoneAltro = u.CanoneAltro
		contratto.CanoneTot = u.CanoneTot
		contratto.KmContrattuali = u.KmContrattuali
		contratto.AddebitoKmExtra = u.AddebitoKmExtra
		contratto.ImportoKm = u.ImportoKm
		contratto.ImportoTot = u.ImportoTot
		contratto.KmInizioGest = u.KmInizioGest
		contratto.KmFineGest = u.KmFineGest
		contratto.Note = u.Note

		//Aggiorno Contratto
		_, err = o.Update(&contratto)
		if err != nil {
			flash.Error("Errore interno nell'aggiornare il contratto")
			flash.Store(&c.Controller)
			return
		}
		//quando modifico il contratto potremi mandare un messaggio direttamente nella finestra dei messaggi nella pagina ViewContrattiNol
		flash.Notice("Contratto di Noleggio Modificato")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)

	}
}

//AddNoleggio permette di caricare un nuovo contratto di noleggio
//Il parametro passato alla pagina è Fornitore.ID
func (c *AutomezziController) AddNoleggio() {
	c.activeContent("automezzi/add/noleggio")
	//verifica se l'utente è loggato
	sess := c.GetSession("portale")
	if sess == nil {
		c.Redirect("/", 302)
		return
	}
	//inizializza i messaggi
	flash := beego.NewFlash()
	//carica la sessione
	m := sess.(map[string]interface{})
	//verifica se si detengono i privilegi per accedere alla pagina
	if m["admin"] != 3 {
		flash.Notice("Non hai i diritti per accedere a questa pagina")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}
	fmt.Printf("hai i diritti")
	//prende i paramentri della pagina nel formato Ncontratto!IDFornitore
	parms := c.Ctx.Input.Param(":parms")
	idforn, err := strconv.Atoi(parms)
	if err != nil {
		fmt.Println(err)
		flash.Error("Impossibile convertire in Int il codice fornitore")
		flash.Store(&c.Controller)
		return
	}
	//Inizializza il contratto
	contratto := models.ContrNoleggi{}

	//Estrae i dati del fornitore partendo dal ID passato come parametro
	fornitore := models.Fornitori{ID: idforn}

	o := orm.NewOrm()
	o.Using("default")
	err = o.Read(&fornitore, "ID")
	if err != nil {
		fmt.Println(err)
		flash.Error("Internal error")
		flash.Store(&c.Controller)
		return
	}
	//Inserisce nel form i dati del fornitore
	defer func(c *AutomezziController, fornitore *models.Fornitori) {

		c.Data["Fornitori"] = fornitore.Descrizione
		c.Data["PIFornitori"] = fornitore.PI

	}(c, &fornitore)

	//XSRF attack defense
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if c.Ctx.Input.Method() == "POST" {

		u := contrattoNol{}
		//Estra i dati dal form
		if err := c.ParseForm(&u); err != nil {
			fmt.Println("cannot parse form")
			return
		}

		//c.Data["data"] = u
		//Valida i dati del form
		valid := validation.Validation{}
		if b, _ := valid.Valid(&u); !b {
			c.Data["Errors"] = valid.ErrorsMap
			return
		}

		//Converte le data da string a time
		dcont, err := time.Parse("2006-01-02 15:04", u.DataCont+" 10:00")
		if err != nil {
			fmt.Println(err)
			return
		}
		diniz, err := time.Parse("2006-01-02 15:04", u.DataInizio+" 10:00")
		if err != nil {
			fmt.Println(err)
			return
		}
		dfine, err := time.Parse("2006-01-02 15:04", u.DataFine+" 10:00")
		if err != nil {
			fmt.Println(err)
			return
		}

		//Inserisce i dati estratti dal form den contratto

		contratto.NContratto = u.NContratto
		contratto.DataCont = dcont
		contratto.DataInizio = diniz
		contratto.DataFine = dfine
		contratto.Riparametrizzazione = u.Riparametrizzazione
		contratto.NRate = u.NRate
		contratto.CanoneBase = u.CanoneBase
		contratto.CanoneServizi = u.CanoneServizi
		contratto.CanoneAltro = u.CanoneAltro
		contratto.CanoneTot = u.CanoneTot
		contratto.KmContrattuali = u.KmContrattuali
		contratto.AddebitoKmExtra = u.AddebitoKmExtra
		contratto.ImportoKm = u.ImportoKm
		contratto.ImportoTot = u.ImportoTot
		contratto.KmInizioGest = u.KmInizioGest
		contratto.KmFineGest = u.KmFineGest
		contratto.Note = u.Note
		contratto.Fornitori = &fornitore

		//inserisce il contratto nel db
		_, err = o.Insert(&contratto)
		if err != nil {
			flash.Error("Errore autorizzazioni applicazioni")
			flash.Store(&c.Controller)
			return
		}

		flash.Notice("Contratto Noleggio Registrato")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)

	}
}

//AddConducente permette di caricare un nuovo conducente
func (c *AutomezziController) AddConducente() {
	c.activeContent("automezzi/add/conducente")

	sess := c.GetSession("portale")
	if sess == nil {
		c.Redirect("/", 302)
		return
	}
	flash := beego.NewFlash()
	m := sess.(map[string]interface{})
	fmt.Println(m["admin"])
	fmt.Println(reflect.ValueOf(m["admin"]).Type())
	if m["admin"] != 3 {
		flash.Notice("Non hai i diritti per accedere a questa pagina")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}
	fmt.Printf("hai i diritti")
	parms := c.Ctx.Input.Param(":parms")
	fmt.Println(parms)
	c.Data["parms"] = parms

	//XSRF attack defense
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if c.Ctx.Input.Method() == "POST" {

		u := conducente{}

		if err := c.ParseForm(&u); err != nil {
			fmt.Println("cannot parse form")
			return
		}
		fmt.Println(u)
		c.Data["Conducente"] = u
		valid := validation.Validation{}
		if b, _ := valid.Valid(&u); !b {
			c.Data["Errors"] = valid.ErrorsMap
			return
		}
		//******** Save user info to database
		o := orm.NewOrm()
		o.Using("default")

		//create user and userApp models

		Conducente := models.Conducenti{Nome: u.Nome, Cognome: u.Cognome, CodiceFiscale: u.CF}

		_, err := o.Insert(&Conducente)
		if err != nil {
			flash.Error("Errore autorizzazioni applicazioni")
			flash.Store(&c.Controller)
			return
		}

		flash.Notice("Conducente Registrato")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}
}

//AddData funzione test percapire il funzionamento dell'inserimento della data
func (c *AutomezziController) AddData() {
	c.activeContent("automezzi/add/appoggio")

	sess := c.GetSession("portale")
	if sess == nil {
		c.Redirect("/", 302)
		return
	}
	flash := beego.NewFlash()
	m := sess.(map[string]interface{})
	fmt.Println(m["admin"])
	fmt.Println(reflect.ValueOf(m["admin"]).Type())
	if m["admin"] != 3 {
		flash.Notice("Non hai i diritti per accedere a questa pagina")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}
	fmt.Printf("hai i diritti")
	parms := c.Ctx.Input.Param(":parms")
	fmt.Println(parms)
	c.Data["parms"] = parms

	//XSRF attack defense
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if c.Ctx.Input.Method() == "POST" {

		u := data1{}

		if err := c.ParseForm(&u); err != nil {
			fmt.Println("cannot parse form")
			return
		}
		fmt.Println("STAMPA DATA")
		fmt.Print(u)
		fmt.Println(u.dataInFlotta)

		//funziona!! ora bisogna trasformarlo in time
		res := c.GetString("dataInFlotta")
		res = res + " 10:00"
		fmt.Println(res)

		t, err := time.Parse("2006-01-02 15:04", res)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(t)
		/*
			valid := validation.Validation{}
			if b, _ := valid.Valid(&u); !b {
				c.Data["Errors"] = valid.ErrorsMap
				return
			}*/
		//******** Save user info to database

		o := orm.NewOrm()
		o.Using("default")

		//create user and userApp models

		Data := models.DateTest{Data: t, Data2: t, Data3: t}

		_, err = o.Insert(&Data)
		if err != nil {
			flash.Error("Errore autorizzazioni applicazioni")
			flash.Store(&c.Controller)
			return
		}
		flash.Notice("Data Registrato")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}
}

//AddAutomezzo permette di aggiungere un nuovo automezzo usando una normale schermata
func (c *AutomezziController) AddAutomezzo() {
	c.activeContent("automezzi/add/automezzo")

	sess := c.GetSession("portale")
	if sess == nil {
		c.Redirect("/", 302)
		return
	}
	flash := beego.NewFlash()
	m := sess.(map[string]interface{})
	fmt.Println(m["admin"])
	fmt.Println(reflect.ValueOf(m["admin"]).Type())
	if m["admin"] != 3 {
		flash.Notice("Non hai i diritti per accedere a questa pagina")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}
	fmt.Printf("hai i diritti")
	parms := c.Ctx.Input.Param(":parms")
	fmt.Println(parms)
	c.Data["parms"] = parms

	var Conducentelist string

	o := orm.NewOrm()
	o.Using("default")
	var conducenti []models.Conducenti

	//conducenti = o.QueryTable("conducenti")

	num, err := o.QueryTable("conducenti").All(&conducenti)

	if err != nil {
		flash.Notice("Errore, contattare l'amministratore del sito")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}

	fmt.Println("Numero Conducenti: ", num)
	for i := range conducenti {

		Conducentelist += fmt.Sprintf("<option value=\"" + strconv.Itoa(conducenti[i].ID) + "\">" + conducenti[i].Nome + " " + conducenti[i].Cognome + "</option>")
	}

	//c.Data["Tipoveicololist"] = template.HTML(Tipoveicololist)
	//c.Data["Condizionelist"] = template.HTML(Condizionelist)
	c.Data["Conducentelist"] = template.HTML(Conducentelist)
	//c.Data["Settorelist"] = template.HTML(Settorelist)
	//c.Data["Impiegolist"] = template.HTML(Impiegolist)
	//c.Data["NormEurolist"] = template.HTML(NormEurolist)

	//XSRF attack defense
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if c.Ctx.Input.Method() == "POST" {

		u := automezzo{}

		if err := c.ParseForm(&u); err != nil {
			fmt.Println("cannot parse form")
			fmt.Println(err)
			return
		}
		fmt.Println(u)
		fmt.Println(u.Targa)
		u.Targa = strings.ToUpper(u.Targa)
		u.NLibretto = strings.ToUpper(u.NLibretto)
		u.NTelaio = strings.ToUpper(u.NTelaio)
		fmt.Println(u)
		fmt.Println(u.Targa)
		c.Data["Automezzo"] = u
		/*
			valid := validation.Validation{}
			if b, _ := valid.Valid(&u); !b {
				c.Data["Errors"] = valid.ErrorsMap
				return
			}
		*/
		//******** Save user info to database
		o := orm.NewOrm()
		o.Using("default")

		//create user and userApp models

		tipoVeicolo := models.TipiVeicolo{ID: u.TipoVeicolo}
		Condizione := models.Condizioni{ID: u.Condizione}
		Conducente := models.Conducenti{ID: u.Conducente}
		Carburante := models.Carburante{ID: u.Carburante}
		Settore := models.Settori{ID: u.Settore}

		impiego := models.Impieghi{ID: u.Impiego}
		fmt.Println(impiego)
		fmt.Println(u.Impiego)

		VeicoloDT := models.VeicoliDT{}

		switch u.Contratto {
		case 0:
			Contratto := models.ContrAcquisti{NContratto: u.NContratto}
			VeicoloDT = models.VeicoliDT{AnnoImmatricolazione: u.AnnoImmatricolazione, NLibretto: u.NLibretto, ContrAcquisti: &Contratto, NTelaio: u.NTelaio, Marca: u.Marca, Modello: u.Modello, Carburante: &Carburante, NorEuro: u.NorEuro, Kw: u.KW, Cilindrata: u.Cilindrata, ConsumoTeorico: u.ConsumoTeorico, KmAnno: u.KMAnno, CostoKm: u.Costokm, Pneumatici: u.Pneumatici}
			o.Insert(&Contratto)

		case 1:

			Contratto := &models.ContrLeasing{NContratto: u.NContratto}
			VeicoloDT = models.VeicoliDT{AnnoImmatricolazione: u.AnnoImmatricolazione, NLibretto: u.NLibretto, NTelaio: u.NTelaio, Marca: u.Marca, Modello: u.Modello, Carburante: &Carburante, NorEuro: u.NorEuro, Kw: u.KW, Cilindrata: u.Cilindrata, ConsumoTeorico: u.ConsumoTeorico, KmAnno: u.KMAnno, CostoKm: u.Costokm, Pneumatici: u.Pneumatici}
			m2m := o.QueryM2M(&VeicoloDT, "ContrLeasing")

			o.Insert(Contratto)

			num, err := m2m.Add(Contratto)
			if err == nil {
				fmt.Println("Contratti Leasing Aggiunti: ", num)
			}

		case 2:

			Contratto := &models.ContrNoleggi{NContratto: u.NContratto}
			VeicoloDT = models.VeicoliDT{AnnoImmatricolazione: u.AnnoImmatricolazione, NLibretto: u.NLibretto, NTelaio: u.NTelaio, Marca: u.Marca, Modello: u.Modello, Carburante: &Carburante, NorEuro: u.NorEuro, Kw: u.KW, Cilindrata: u.Cilindrata, ConsumoTeorico: u.ConsumoTeorico, KmAnno: u.KMAnno, CostoKm: u.Costokm, Pneumatici: u.Pneumatici}
			m2m := o.QueryM2M(&VeicoloDT, "ContrNoleggi")
			o.Insert(Contratto)

			num, err := m2m.Add(Contratto)
			if err == nil {
				fmt.Println("Contratti Noleggio Aggiunti: ", num)
			}
		}

		VeicoloDG := models.VeicoliDG{Targa: u.Targa, DataInFlotta: u.DataInFlotta, DataFineFlotta: u.DataFineFlotta, TipiVeicolo: &tipoVeicolo, Settori: &Settore, Condizioni: &Condizione, Impieghi: &impiego, Conducenti: &Conducente, Note: u.Note, VeicoliDT: &VeicoloDT}

		_, err = o.Insert(&VeicoloDT)
		if err != nil {
			fmt.Println(err)
			flash.Error("Errore autorizzazioni applicazioni - VeicoloDT")
			flash.Store(&c.Controller)
			return
		}

		_, err = o.Insert(&VeicoloDG)
		if err != nil {
			fmt.Println(err)
			flash.Error("Errore autorizzazioni applicazioni - impossibile registrare veicolo")
			flash.Store(&c.Controller)
			return
		}

		flash.Notice("Veicolo Registrato")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}

}

//AddAutomezzo2 permette di aggiungere un nuovo automezzo usando una schermata con i tab
func (c *AutomezziController) AddAutomezzo2() {
	c.activeContent("automezzi/add/automezzo2")

	sess := c.GetSession("portale")
	if sess == nil {
		c.Redirect("/", 302)
		return
	}
	flash := beego.NewFlash()
	m := sess.(map[string]interface{})
	fmt.Println(m["admin"])
	fmt.Println(reflect.ValueOf(m["admin"]).Type())
	if m["admin"] != 3 {
		flash.Notice("Non hai i diritti per accedere a questa pagina")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}
	fmt.Printf("hai i diritti")
	parms := c.Ctx.Input.Param(":parms")
	fmt.Println(parms)
	c.Data["parms"] = parms

	var Conducentelist string

	o := orm.NewOrm()
	o.Using("default")
	var conducenti []models.Conducenti

	//conducenti = o.QueryTable("conducenti")

	num, err := o.QueryTable("conducenti").All(&conducenti)

	if err != nil {
		flash.Notice("Errore, contattare l'amministratore del sito")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}

	fmt.Println("Numero Conducenti: ", num)
	for i := range conducenti {

		Conducentelist += fmt.Sprintf("<option value=\"" + strconv.Itoa(conducenti[i].ID) + "\">" + conducenti[i].Nome + " " + conducenti[i].Cognome + "</option>")
	}

	c.Data["Conducentelist"] = template.HTML(Conducentelist)

	//XSRF attack defense
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if c.Ctx.Input.Method() == "POST" {

		u := automezzo{}

		if err := c.ParseForm(&u); err != nil {
			fmt.Println("cannot parse form")
			fmt.Println(err)
			return
		}

		u.Targa = strings.ToUpper(u.Targa)
		u.NLibretto = strings.ToUpper(u.NLibretto)
		u.NTelaio = strings.ToUpper(u.NTelaio)
		c.Data["Automezzo"] = u

		valid := validation.Validation{}
		if b, _ := valid.Valid(&u); !b {
			c.Data["Errors"] = valid.ErrorsMap
			return
		}

		//******** Save user info to database
		o := orm.NewOrm()
		o.Using("default")

		//create user and userApp models

		tipoVeicolo := models.TipiVeicolo{ID: u.TipoVeicolo}
		Condizione := models.Condizioni{ID: u.Condizione}
		Conducente := models.Conducenti{ID: u.Conducente}
		Carburante := models.Carburante{ID: u.Carburante}
		Settore := models.Settori{ID: u.Settore}

		impiego := models.Impieghi{ID: u.Impiego}

		VeicoloDT := models.VeicoliDT{}

		switch u.Contratto {
		case 0:

			Contratto := models.ContrAcquisti{NContratto: u.NContratto}
			fmt.Println("Contratto :")
			fmt.Println(Contratto)
			_, err = o.Insert(&Contratto)
			if err != nil {
				fmt.Println(err)
				flash.Error("Errore Inserimento Contratto")
				flash.Store(&c.Controller)
				return
			}
			fmt.Println(err)
			VeicoloDT = models.VeicoliDT{AnnoImmatricolazione: u.AnnoImmatricolazione, NLibretto: u.NLibretto, ContrAcquisti: &Contratto, NTelaio: u.NTelaio, Marca: u.Marca, Modello: u.Modello, Carburante: &Carburante, NorEuro: u.NorEuro, Kw: u.KW, Cilindrata: u.Cilindrata, ConsumoTeorico: u.ConsumoTeorico, KmAnno: u.KMAnno, CostoKm: u.Costokm, Pneumatici: u.Pneumatici}
			fmt.Println(Contratto)
			fmt.Println(VeicoloDT)
		case 1:

			Contratto := models.ContrAcquisti{NContratto: u.NContratto}
			fmt.Println("Contratto :")
			fmt.Println(Contratto)
			_, err = o.Insert(&Contratto)
			if err != nil {
				fmt.Println(err)
				flash.Error("Errore Inserimento Contratto")
				flash.Store(&c.Controller)
				return
			}
			fmt.Println(err)
			VeicoloDT = models.VeicoliDT{AnnoImmatricolazione: u.AnnoImmatricolazione, NLibretto: u.NLibretto, NTelaio: u.NTelaio, Marca: u.Marca, Modello: u.Modello, Carburante: &Carburante, NorEuro: u.NorEuro, Kw: u.KW, Cilindrata: u.Cilindrata, ConsumoTeorico: u.ConsumoTeorico, KmAnno: u.KMAnno, CostoKm: u.Costokm, Pneumatici: u.Pneumatici}
			m2m := o.QueryM2M(&VeicoloDT, "ContrLeasing")

			o.Insert(Contratto)

			num, err := m2m.Add(Contratto)
			if err == nil {
				fmt.Println("Contratti Leasing Aggiunti: ", num)
			}

		case 2:

			Contratto := models.ContrAcquisti{NContratto: u.NContratto}
			fmt.Println("Contratto :")
			fmt.Println(Contratto)
			_, err = o.Insert(&Contratto)
			if err != nil {
				fmt.Println(err)
				flash.Error("Errore Inserimento Contratto")
				flash.Store(&c.Controller)
				return
			}
			fmt.Println(err)
			VeicoloDT = models.VeicoliDT{AnnoImmatricolazione: u.AnnoImmatricolazione, NLibretto: u.NLibretto, NTelaio: u.NTelaio, Marca: u.Marca, Modello: u.Modello, Carburante: &Carburante, NorEuro: u.NorEuro, Kw: u.KW, Cilindrata: u.Cilindrata, ConsumoTeorico: u.ConsumoTeorico, KmAnno: u.KMAnno, CostoKm: u.Costokm, Pneumatici: u.Pneumatici}
			m2m := o.QueryM2M(&VeicoloDT, "ContrNoleggi")
			o.Insert(Contratto)

			num, err := m2m.Add(Contratto)
			if err == nil {
				fmt.Println("Contratti Noleggio Aggiunti: ", num)
			}
		}

		VeicoloDG := models.VeicoliDG{Targa: u.Targa, DataInFlotta: u.DataInFlotta, DataFineFlotta: u.DataFineFlotta, TipiVeicolo: &tipoVeicolo, Settori: &Settore, Condizioni: &Condizione, Impieghi: &impiego, Conducenti: &Conducente, Note: u.Note, VeicoliDT: &VeicoloDT}

		_, err = o.Insert(&VeicoloDT)
		if err != nil {
			fmt.Println(err)
			flash.Error("Errore autorizzazioni applicazioni - VeicoloDT")
			flash.Store(&c.Controller)
			return
		}

		_, err = o.Insert(&VeicoloDG)
		if err != nil {
			fmt.Println(err)
			flash.Error("Errore autorizzazioni applicazioni - impossibile registrare veicolo")
			flash.Store(&c.Controller)
			return
		}

		flash.Notice("Veicolo Registrato")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}

}

//ViewAutomezzi permette di visualizzare gli automezzi caricati sul portale
func (c *AutomezziController) ViewAutomezzi() {
	// Only administrator can Manage accounts
	c.activeContent("automezzi/view/automezzi")

	//******** c page requires login
	sess := c.GetSession("portale")
	if sess == nil {
		c.Redirect("/", 302)
		return
	}
	flash := beego.NewFlash()
	m := sess.(map[string]interface{})
	fmt.Println(m["admin"])
	fmt.Println(reflect.ValueOf(m["admin"]).Type())
	if m["admin"] != 3 {
		flash.Notice("Non hai i diritti per accedere a questa pagina")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}

	fmt.Printf("hai i diritti")

	//in caso di panic reindirizza alla home
	defer func(c *AutomezziController) {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Index", r)
			c.Redirect("/", 302)
		}
	}(c)
	//NON VA SENZA PARAMETRI
	//parametro serve per suddividere in pagine l'elenco contratti e se non è presente viene forzato
	const pagesize = 10
	var parms string
	if c.Ctx.Input.Param(":parms") == "" {
		parms = "id!0!id__gte,0"
	} else {
		parms = c.Ctx.Input.Param(":parms")
	}

	c.Data["parms"] = parms
	str := strings.Split(parms, "!")
	fmt.Println("parms is", str)
	order := str[0]
	off, _ := strconv.Atoi(str[1])
	offset := int64(off)
	if offset < 0 {
		offset = 0
	}
	query := str[2]

	rows := ""
	var veicoli []*models.VeicoliDG
	qs, ok := c.setCompare(query, "veicoli_d_g")
	if !ok {
		fmt.Println("cannot set QuerySeter")
		o := orm.NewOrm()
		qs := o.QueryTable("veicoli_d_g")
		qs.RelatedSel("veicoli_d_t")
		qs = qs.Filter("id__gte", 0)
		c.Data["query"] = "id__gte,0"
	}

	count, _ := qs.Count()
	c.Data["count"] = count
	if offset >= count {
		offset = 0
	}

	num, err := qs.Limit(pagesize, offset).OrderBy(order).RelatedSel().All(&veicoli)
	if err != nil {
		fmt.Println("Query table failed:", err)
	}

	/*TABELLA IN BASE AI PARAMETRI*/
	for i := range veicoli {
		rows += fmt.Sprintf("<tr><td>%s</td>"+
			"<td>%s</td><td>%s</td><td><a class='btn btn-defalut btn-xs' href='http://%s/automezzi/edit/veicoli/%d'><span class='glyphicon glyphicon-edit'></span> Modifica</a></td></tr>", veicoli[i].Targa, veicoli[i].VeicoliDT.Marca, veicoli[i].VeicoliDT.Modello, appcfgdomainname, veicoli[i].ID)
	}
	c.Data["Rows"] = template.HTML(rows)

	c.Data["order"] = order
	c.Data["offset"] = offset
	c.Data["end"] = max(0, (count/pagesize)*pagesize)

	if num+offset < count {
		c.Data["next"] = num + offset
	}
	if offset-pagesize >= 0 {
		c.Data["prev"] = offset - pagesize
		c.Data["showprev"] = true
	} else if offset > 0 && offset < pagesize {
		c.Data["prev"] = 0
		c.Data["showprev"] = true
	}

	if count > pagesize {
		c.Data["ShowNav"] = true
	}
	c.Data["progress"] = float64(offset*100) / float64(max(count, 1))

}

//EditAutomezzo permette di modificare gli automezzi caricati sul portale
//Il parametro passato alla pagina è VeicoliDG.ID
func (c *AutomezziController) EditAutomezzo() {
	/*
		c.activeContent("automezzi/edit/automezzo")
		//verifica se l'utente è loggato
		sess := c.GetSession("portale")
		if sess == nil {
			c.Redirect("/", 302)
			return
		}
		//inizializza i messaggi
		flash := beego.NewFlash()
		//carica la sessione
		m := sess.(map[string]interface{})
		//verifica se si detengono i privilegi per accedere alla pagina
		if m["admin"] != 3 {
			flash.Notice("Non hai i diritti per accedere a questa pagina")
			flash.Store(&c.Controller)
			c.Redirect("/notice", 302)
		}
		fmt.Printf("hai i diritti")
		//prende i paramentri della pagina nel formato Ncontratto!IDFornitore
		parms := c.Ctx.Input.Param(":parms")
		idveic, _ := strconv.Atoi(parms)
		c.Data["AutomezzoDG.ID"] = idveic
		veicolo := models.VeicoliDG{ID: idveic}

		o := orm.NewOrm()
		o.Using("default")
		//Flag per stabilire se il contratto è nuovo o esisteva già

		err := o.QueryTable("veicoli_d_g").Filter("ID", idveic).RelatedSel().One(&veicolo)
		if err == orm.ErrNoRows {
			//non esiste nessun contratto con questo codice
			flash.Notice("Il veicolo non esiste")
			flash.Store(&c.Controller)
			c.Redirect("/notice", 302)
		} else if err == orm.ErrMissPK {
			fmt.Println("Errore - Contattare l'amministratore del sito")
		}
		//aggiorna il form con i dati letti dal DB

		defer func(c *AutomezziController, veicolo *models.VeicoliDG) {

			c.Data["AutomezzoDG.ID"] = veicolo.ID
			c.Data["AutomezzoDG.Targa"] = veicolo.Targa
			c.Data["AutomezzoDG.DataInFlotta"] = veicolo.DataInFlotta
			c.Data["AutomezzoDG.DataFineFlotta"] = veicolo.DataFineFlotta
			c.Data["AutomezzoDG.Note"] = veicolo.Note
			c.Data["AutomezzoDG.TipoVeicolo"] = veicolo.TipiVeicolo.Descrizione
			c.Data["AutomezzoDG.Settore"] = veicolo.Settori.Descrizione
			c.Data["AutomezzoDG.Condizioni"] = veicolo.Condizioni
			c.Data["AutomezzoDG.Impiegho"] = veicolo.Impieghi.Descrizione
			c.Data["AutomezzoDG.Conducenti"] = veicolo.Conducenti // bisogna gestire i multi conducenti
			c.Data["AutomezzoDG.VeicoliDT.AnnoImmatricolazione"] = veicolo.VeicoliDT.AnnoImmatricolazione
			c.Data["AutomezzoDG.VeicoliDT.NLibretto"] = veicolo.VeicoliDT.NLibretto
			c.Data["AutomezzoDG.VeicoliDT.NTelaio"] = veicolo.VeicoliDT.NTelaio
			c.Data["AutomezzoDG.VeicoliDT.Marca"] = veicolo.VeicoliDT.Marca
			c.Data["AutomezzoDG.VeicoliDT.Modello"] = veicolo.VeicoliDT.Modello
			c.Data["AutomezzoDG.VeicoliDT.NorEuro"] = veicolo.VeicoliDT.NorEuro
			c.Data["AutomezzoDG.VeicoliDT.Kw"] = veicolo.VeicoliDT.Kw
			c.Data["AutomezzoDG.VeicoliDT.Cilindrata"] = veicolo.VeicoliDT.Cilindrata
			c.Data["AutomezzoDG.VeicoliDT.ConsumoTeorico"] = veicolo.VeicoliDT.ConsumoTeorico

			c.Data["AutomezzoDG.VeicoliDT.KmAnno"] = veicolo.VeicoliDT.KmAnno
			c.Data["AutomezzoDG.VeicoliDT.CostoKm"] = veicolo.VeicoliDT.CostoKm
			c.Data["AutomezzoDG.VeicoliDT.Pneumatici"] = veicolo.VeicoliDT.Pneumatici

			c.Data["AutomezzoDG.VeicoliDT.Allegati"] = veicolo.VeicoliDT.Allegati //bisogna gestire i multi allegati
			c.Data["AutomezzoDG.VeicoliDT.Carburante"] = veicolo.VeicoliDT.Carburante.Descrizione
			c.Data["AutomezzoDG.VeicoliDT.ContrAcquisti"] = veicolo.VeicoliDT.ContrAcquisti
			c.Data["AutomezzoDG.VeicoliDT.ContrLeasing"] = veicolo.VeicoliDT.ContrLeasing
			c.Data["AutomezzoDG.VeicoliDT.ContrNoleggi"] = veicolo.VeicoliDT.ContrNoleggi //bisogna gestire i multi contratti

		}(c, &veicolo)

		//XSRF attack defense
		c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
		if c.Ctx.Input.Method() == "POST" {

			u := automezzo{}

			if err := c.ParseForm(&u); err != nil {
				fmt.Println("cannot parse form")
				return
			}
			f := models.Fornitori{ID: contratto.Fornitori.ID, Descrizione: u.Fornitori, PI: u.PIFornitori}
			fmt.Println(u.NContratto)
			fmt.Println(u.DataCont)
			fmt.Println(f.Descrizione)
			c.Data["data"] = u
			valid := validation.Validation{}
			if b, _ := valid.Valid(&u); !b {
				c.Data["Errors"] = valid.ErrorsMap
				return
			}
			//Aggiorno Fornitore
			_, err := o.Update(&f)
			if err != nil {
				flash.Error("Errore interno")
				flash.Store(&c.Controller)
				return
			}
			//Converte le data da string a time
			dcont, err := time.Parse("2006-01-02 15:04", u.DataCont+" 10:00")
			if err != nil {
				fmt.Println(err)
				return
			}
			diniz, err := time.Parse("2006-01-02 15:04", u.DataInizio+" 10:00")
			if err != nil {
				fmt.Println(err)
				return
			}
			dfine, err := time.Parse("2006-01-02 15:04", u.DataFine+" 10:00")
			if err != nil {
				fmt.Println(err)
				return
			}

			veicolo.Targa = u.Targa
			veicolo.DataInFlotta = u.DataInFlotta
			veicolo.DataFineFlotta = u.DataFineFlotta
			veicolo.TipoVeicolo = u.TipoVeicolo
			veicolo.Settore = u.Settore
			veicolo.Condizione = u.Condizione
			veicolo.Impiego = u.Impiego
			veicolo.Conducente = u.Conducente
			veicolo.Note = u.Note
			veicolo.AnnoImmatricolazione = u.AnnoImmatricolazione
			veicolo.NLibretto = u.NLibretto
			veicolo.NTelaio = u.NTelaio
			veicolo.Marca = u.Marca
			veicolo.Modello = u.Modello
			veicolo.Carburante = u.Carburante
			veicolo.NorEuro = u.NorEuro
			veicolo.KW = u.KW
			veicolo.Cilindrata = u.Cilindrata
			veicolo.ConsumoTeorico = u.ConsumoTeorico
			veicolo.KMAnno = u.KMAnno
			veicolo.Costokm = u.Costokm
			veicolo.Pneumatici = u.Pneumatici
			veicolo.Contratto = u.Contratto
			veicolo.NContratto = u.NContratto

			//Aggiorno Contratto
			_, err = o.Update(&veicolo)
			if err != nil {
				flash.Error("Errore interno nell'aggiornare il contratto")
				flash.Store(&c.Controller)
				return
			}
			//quando modifico il contratto potremi mandare un messaggio direttamente nella finestra dei messaggi nella pagina ViewContrattiNol
			flash.Notice("Contratto di Noleggio Modificato")
			flash.Store(&c.Controller)
			c.Redirect("/notice", 302)

		}
	*/
}

//AddAutomezzo3 permette di caricare un nuovo automezzo
//Il parametro passato alla pagina è Fornitore.ID e ID Contratti
func (c *AutomezziController) AddAutomezzo3() {
	c.activeContent("automezzi/add/automezzo")
	//verifica se l'utente è loggato
	sess := c.GetSession("portale")
	if sess == nil {
		c.Redirect("/", 302)
		return
	}
	//inizializza i messaggi
	flash := beego.NewFlash()
	//carica la sessione
	m := sess.(map[string]interface{})
	//verifica se si detengono i privilegi per accedere alla pagina
	if m["admin"] != 3 {
		flash.Notice("Non hai i diritti per accedere a questa pagina")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}
	fmt.Printf("hai i diritti")
	//prende i paramentri della pagina nel formato Ncontratto!IDFornitore
	parms := c.Ctx.Input.Param(":parms")
	idcontr, err := strconv.Atoi(parms)
	if err != nil {
		fmt.Println(err)
		flash.Error("Impossibile convertire in Int il codice fornitore")
		flash.Store(&c.Controller)
		return
	}
	//Inizializza il contratto
	contratto := models.ContrNoleggi{ID: idcontr}

	//Estrae i dati del fornitore partendo dal ID passato come parametro
	fornitore := models.Fornitori{}

	o := orm.NewOrm()
	o.Using("default")
	err = o.Read(&fornitore, "ID")
	if err != nil {
		fmt.Println(err)
		flash.Error("Internal error")
		flash.Store(&c.Controller)
		return
	}
	//Inserisce nel form i dati del fornitore
	defer func(c *AutomezziController, fornitore *models.Fornitori) {

		c.Data["Fornitori"] = fornitore.Descrizione
		c.Data["PIFornitori"] = fornitore.PI

	}(c, &fornitore)

	//XSRF attack defense
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if c.Ctx.Input.Method() == "POST" {

		u := contrattoNol{}
		//Estra i dati dal form
		if err := c.ParseForm(&u); err != nil {
			fmt.Println("cannot parse form")
			return
		}

		//c.Data["data"] = u
		//Valida i dati del form
		valid := validation.Validation{}
		if b, _ := valid.Valid(&u); !b {
			c.Data["Errors"] = valid.ErrorsMap
			return
		}

		//Converte le data da string a time
		dcont, err := time.Parse("2006-01-02 15:04", u.DataCont+" 10:00")
		if err != nil {
			fmt.Println(err)
			return
		}
		diniz, err := time.Parse("2006-01-02 15:04", u.DataInizio+" 10:00")
		if err != nil {
			fmt.Println(err)
			return
		}
		dfine, err := time.Parse("2006-01-02 15:04", u.DataFine+" 10:00")
		if err != nil {
			fmt.Println(err)
			return
		}

		//Inserisce i dati estratti dal form den contratto

		contratto.NContratto = u.NContratto
		contratto.DataCont = dcont
		contratto.DataInizio = diniz
		contratto.DataFine = dfine
		contratto.Riparametrizzazione = u.Riparametrizzazione
		contratto.NRate = u.NRate
		contratto.CanoneBase = u.CanoneBase
		contratto.CanoneServizi = u.CanoneServizi
		contratto.CanoneAltro = u.CanoneAltro
		contratto.CanoneTot = u.CanoneTot
		contratto.KmContrattuali = u.KmContrattuali
		contratto.AddebitoKmExtra = u.AddebitoKmExtra
		contratto.ImportoKm = u.ImportoKm
		contratto.ImportoTot = u.ImportoTot
		contratto.KmInizioGest = u.KmInizioGest
		contratto.KmFineGest = u.KmFineGest
		contratto.Note = u.Note
		contratto.Fornitori = &fornitore

		//inserisce il contratto nel db
		_, err = o.Insert(&contratto)
		if err != nil {
			flash.Error("Errore autorizzazioni applicazioni")
			flash.Store(&c.Controller)
			return
		}

		flash.Notice("Contratto Noleggio Registrato")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)

	}
}
