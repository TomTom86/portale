<div class="col-md-10 content">
    <div class="panel panel-default">
        <div class="panel-heading">
            <h3>Modifica Veicolo</h3>
        </div>
        <div class="panel-body">
            {{if .flash.error}}
            <div class="alert alert-danger">
                <button type="button" class="close" data-dismiss="alert" aria-hidden="true">
                    ×</button>
                <span class="glyphicon glyphicon-hand-right"></span> <strong> ATTENZIONE - Errore1</strong>
                <hr class="message-inner-separator">
                {{.flash.error}}.
            </div>
            {{end}}
            {{if .flash.notice}}
            <div class="alert alert-success">
                <button type="button" class="close" data-dismiss="alert" aria-hidden="true">
                    ×</button>
               <span class="glyphicon glyphicon-ok"></span> <strong>Operazione completata con successo</strong>
                <hr class="message-inner-separator">
                {{.flash.notice}}.
            </div>
            {{end}}            
            {{if .Errors}}
            {{range $rec := .Errors}}
            <div class="alert alert-danger">
                <button type="button" class="close" data-dismiss="alert" aria-hidden="true">
                    ×</button>
                <span class="glyphicon glyphicon-hand-right"></span> <strong> ATTENZIONE - Errore</strong>
                <hr class="message-inner-separator">
                {{$rec}}.
            </div>      
            {{end}}
            &nbsp;
            {{end}} 
                <form class="form-horizontal" method="POST">
                    {{ .xsrfdata }}         
                    <!-- TARGA --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="first">Targa</label>                          
                        <div class="col-md-4">
                        {{if .Errors.Targa}}  {{.Errors.Targa}}  {{end}}
                        <input id="targa" name="targa" type="text" value="{{.AutomezzoDG.Targa}}" size="7" class="form-control input-md"/>
                        </div> 
                    </div> 
                    <!-- DATA IN FLOTTA --> 
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="dataInFlotta">Data entrata in flotta</label>                     
                        <div class="col-md-4">
                        {{if .Errors.AutomezzoDG.DataInFlotta}}  {{.Errors.AutomezzoDG.DataInFlotta}}  {{end}}
                        <input id="dataInFlotta" name="dataInFlotta" type=date min=2000-01-01 datetime="2000-01-01 00:00:01" value="{{.AutomezzoDG.DataInFlotta}}" class="form-control input-md"/>
                        </div> 
                    </div>

                    <!-- DATA FINE FLOTTA -->            
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="dataFineFlotta">Data uscita da flotta</label>
                        <div class="col-md-4">
                        {{if .Errors.AutomezzoDG.DataFineFlotta}}  {{.Errors.AutomezzoDG.DataFineFlotta}}  {{end}}
                        <input id="dataFineFlotta" name="dataFineFlotta" type=date min=2000-01-01 datetime="2000-01-01 00:00:01" value="{{.AutomezzoDG.DataFineFlotta}}" size="10" class="form-control input-md"/>
                        </div>
                    </div>   
                    <!-- TIPOVEICOLO --> 
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="tipoteicolo">Tipo Veicolo:</label>
                        <div class="col-md-4">  
                        <select id="tipoveicolo" name="tipoveicolo" type="number" >
                            	<option value="1" selected="selected">Autoveicolo</option>
                                <option value="2">Camion</option>
                                <option value="3">Ciclomotore Stato</option>
                                <option value="4">Furgone</option>
                                <option value="5">Pullman</option>
                                <option value="6">Motoveicolo</option>
                                <option value="7">Altro</option>
                        </select>
                        </div>
                    </div>  
                    <!-- CONDIZIONE --> 
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="condizione">Condizione:</label>
                        <div class="col-md-4">  
                        <select id="condizione" name="condizione" type="number" size="1">
                            <option value="1" selected="selected">Buono stato</option>
                            <option value="2">Cattivo stato</option>
                            <option value="3">Discreto Stato</option>
                            <option value="4">In attesa di alienzazione</option>
                            <option value="5">In attesa di Assegnazione</option>
                            <option value="6">In attesa di Riparazione</option>
                            <option value="7">Non utilizzabile</option>
                            <option value="8">Rubato</option>
                            <option value="9">Alienato</option>
                        </select>
                        </div>
                    </div>   
                    <!-- CONDUCENTE --> 
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="conducente">Conducente:</label>
                        <div class="col-md-4">    
                        <select id="conducente" name="conducente" type="number" size="1">
                            {{.Conducentelist}}
                        </select>
                        </div>
                    </div>  
                    <!-- SETTORE --> 
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="settore">Settore:</label>
                        <div class="col-md-4">  
                        <select id="usettore" name="settore" type="number" size="1">
                            <option value="1">Food</option>
                            <option value="2">Lavanderia</option>
                            <option value="3">Pulizia</option>
                            <option value="4">Marketing</option>
                            <option value="5">Officina</option>
                            <option value="6" selected="selected">Agenti</option>
                            <option value="7">Direzione</option>
                        </select>
                        </div>
                    </div>  
                    <!-- IMPIEGO --> 
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="impiego">Impiego:</label>
                        <div class="col-md-4">  
                        <select id="impiego" name="impiego" type="number" size="1">
                            <option value="1" selected="selected">Aziendale</option>
                            <option value="2">Aziendale + Personale</option>
                            <option value="3">Personale</option>
                        </select>
                        </div>
                    </div>    
                    <!-- NOTE --> 
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="note">Note</label> 
                        <div class="col-md-4"> 
                        {{if .Errors.AutomezzoDG.Note}}  {{.Errors.AutomezzoDG.Note}}  {{end}}    
                            <input id="note" name="note" type="text" value="{{.AutomezzoDG.Note}}" size="100" class="form-control input-md"/>
                        </div>
                    </div>   
                    <!-- ANNO IMMATRICOLAZIONE -->               
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="annoImmatricolazione">Anno immatricolazione</label>  
                        <div class="col-md-4"> 
                        {{if .Errors.AutomezzoDT.AnnoImmatricolazione}}  {{.Errors.AutomezzoDT.AnnoImmatricolazione}}  {{end}}
                            <input id="annoImmatricolazione" name="annoImmatricolazione" type="number" min="1950" value="{{.AutomezzoDT.AnnoImmatricolazione}}" size="4" class="form-control input-md"/>
                        </div>
                    </div>  
                    <!-- TIPO CONTRATTO -->               
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="contratto">Tipo Contratto</label>  
                        <div class="col-md-4"> 
                        {{if .Errors.Contratto}}  {{.Errors.Contratto}}  {{end}}
                        <select id="contratto" name="contratto" size="1">
                            	<option value="0"selected="selected">Acquisto</option>
                                <option value="1">Leasing</option>
                                <option value="2">Noleggio</option>
                        </select>
                        </div>
                    </div> 
                    <!-- NUMERO CONTRATTO -->               
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="numeroContratto">Numero Contratto</label>  
                        <div class="col-md-4"> 
                        {{if .Errors.NContratto}}  {{.Errors.NContratto}}  {{end}}
                            <input id="numeroContratto" name="numeroContratto" type="text" value="{{.NContratto}}" size="20" class="form-control input-md"/>
                        </div>
                    </div>                     
                    <!-- NUMERO LIBRETTO -->               
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="numeroLibretto">Numero libretto</label>  
                        <div class="col-md-4"> 
                        {{if .Errors.AutomezzoDT.NLibretto}}  {{.Errors.AutomezzoDT.NLibretto}}  {{end}}
                            <input id="numeroLibretto" name="numeroLibretto" type="text" value="{{.AutomezzoDT.NLibretto}}" size="20" class="form-control input-md"/>
                        </div>
                    </div>  
                    <!-- NUMERO TELAIO -->               
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="numeroTelaio">Numero di telaio</label>  
                        <div class="col-md-4"> 
                        {{if .Errors.AutomezzoDT.NLibretto}}  {{.Errors.AutomezzoDT.NLibretto}}  {{end}}
                            <input id="numeroTelaio" name="numeroTelaio" type="text" value="{{.AutomezzoDT.NLibretto}}" size="17" class="form-control input-md"/>
                        </div>
                    </div>  
                    <!-- MARCA -->               
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="marca">Marca</label>  
                        <div class="col-md-4"> 
                        {{if .Errors.AutomezzoDT.Marca}}  {{.Errors.AutomezzoDT.Marca}}  {{end}}
                            <input id="marca" name="marca" type="text" value="{{.AutomezzoDT.Marca}}" size="15" class="form-control input-md"/>
                        </div>
                    </div>  
                    <!-- MODELLO -->               
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="modello">Modello</label>  
                        <div class="col-md-4"> 
                        {{if .Errors.AutomezzoDT.Modello}}  {{.Errors.AutomezzoDT.Modello}}  {{end}}
                            <input id="modello" name="modello" type="text" value="{{.AutomezzoDT.Modello}}" size="15" class="form-control input-md"/>
                        </div>
                    </div>  
                    <!-- NORMATIVA EURO -->               
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="normativaEuro">Normativa Euro:</label>
                        <div class="col-md-4">  
                        {{if .Errors.First}}  {{.Errors.First}}  {{end}}  
                        <select id="normativaEuro" name="normativaEuro" size="1">
                            	<option value="0">Euro 0</option>
                                <option value="1">Euro 1</option>
                                <option value="2">Euro 2</option>
                                <option value="3">Euro 3</option>
                                <option value="4">Euro 4</option>
                                <option value="5" selected="selected">Euro 5</option>
                                <option value="6">Euro 6</option>
                        </select>
                        </div>
                    </div> 
                    <!-- KW-->               
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="kw">Kw</label>  
                        <div class="col-md-4"> 
                        {{if .Errors.AutomezzoDT.Kw}}  {{.Errors.AutomezzoDT.Kw}}  {{end}}
                            <input id="kw" name="kw" type="number" min="0" value="{{.AutomezzoDT.Kw}}" size="4" class="form-control input-md"/>
                        </div>
                    </div>                     
                    <!-- CILINDRATA-->               
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="cilindrata">Cilindrata</label>  
                        <div class="col-md-4"> 
                        {{if .Errors.AutomezzoDT.Cilindrata}}  {{.Errors.AutomezzoDT.Cilindrata}}  {{end}}
                            <input id="cilindrata" name="cilindrata" type="number" min="0" value="{{.AutomezzoDT.Cilindrata}}" size="6" class="form-control input-md"/>
                        </div>
                    </div>  
                    <!-- CARBURANTE-->               
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="cilindrata">Carburante</label>  
                        <div class="col-md-4"> 
                        {{if .Errors.AutomezzoDT.Carburante}}  {{.Errors.AutomezzoDT.Carburante}}  {{end}}
                        <select id="carburante" name="carburante" type="number" size="1">
                            <option value="1" selected="selected">Benzina</option>
                            <option value="2">Disel</option>
                            <option value="3">Gas</option>                            
                            <option value="4">Metano</option>
                        </select>
                        </div>
                    </div>                     
                    <!-- CONSUMO TEORICO -->               
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="consumoTeorico">Consumo Teorico</label>  
                        <div class="col-md-4"> 
                        {{if .Errors.AutomezzoDT.ConsumoTeorico}}  {{.Errors.AutomezzoDT.ConsumoTeorico}}  {{end}}
                            <input id="consumoTeorico" name="consumoTeorico" min="0" value="{{.AutomezzoDT.ConsumoTeorico}}" size="4" class="form-control input-md"/>
                        </div>
                    </div>  
                    <!-- KM ANNUI -->               
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="kmAnno">Km Annui</label>  
                        <div class="col-md-4"> 
                        {{if .Errors.AutomezzoDT.KmAnno}}  {{.Errors.AutomezzoDT.KmAnno}}  {{end}}
                            <input id="kmAnno" name="kmAnno" type="number" min="0" value="{{.AutomezzoDT.KmAnno}}" size="10" class="form-control input-md"/>
                        </div>
                    </div>  
                    <!-- COSTO KM -->               
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="costoKm">Costo per km</label>  
                        <div class="col-md-4"> 
                        {{if .Errors.AutomezzoDT.CostoKm}}  {{.Errors.AutomezzoDT.CostoKm}}  {{end}}
                            <input id="costoKm" name="costoKm" type="number" min="0" value="{{.AutomezzoDT.CostoKm}}" size="15" class="form-control input-md" pattern="^\\$?(([1-9](\\d*|\\d{0,2}(,\\d{3})*))|0)(\\.\\d{1,2})?$"/>
                        </div>
                    </div>  
                    <!-- PNEUMATICI -->               
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="pneumatici">Pneumatici</label>  
                        <div class="col-md-4"> 
                        {{if .Errors.AutomezzoDT.Pneumatici}}  {{.Errors.AutomezzoDT.Pneumatici}}  {{end}}
                            <input id="pneumatici" name="pneumatici" type="text "value="{{.AutomezzoDT.Pneumatici}}" size="15" class="form-control input-md"/>
                        </div>
                    </div>                                                                                                                                                                                      
                    <!-- footer --> 
                    <div class="col-md-12 form-group">
                        <button class="btn btn-primary pull-right" type="submit" class="pull-right" value="Inserisci" />Inserisci</button>
                                  
                    </div>                              

                </form>
        </div>
    </div>
</div>




















