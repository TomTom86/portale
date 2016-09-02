

<div class="col-md-10 content">
    <div class="panel panel-default">
        <div class="panel-heading">
            <h3>Modifica Contratto Noleggio</h3>
        </div>
        <div class="panel-body">
            {{if .flash.error}}
            <div class="alert alert-danger">
                <button type="button" class="close" data-dismiss="alert" aria-hidden="true">
                    ×</button>
                <span class="glyphicon glyphicon-hand-right"></span> <strong> ATTENZIONE - Errore</strong>
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
                            
                    <!-- NUMERO CONTRATTO --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Numero Contratto</label>                          
                        <div class="col-md-4">
                        <input id="ncontratto" name="ncontratto" type="text" value="{{.NContratto}}" class="form-control input-md"/>
                        </div> 
                    </div> 
                    <!-- DATA CONTRATTO --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Data Contratto</label>                          
                        <div class="col-md-4">
                        <input id="DataCont" name="DataCont" type=date min=2000-01-01 datetime="2000-01-01 00:00:01" class="form-control input-md" value="{{.DataCont}}" />
                        </div> 
                    </div> 
                    <!-- DATA INIZIO CONTRATTO --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Data Inizio Contratto</label>                          
                        <div class="col-md-4">
                        <input id="datainiziocontr" name="datainiziocontr" type=date min=2000-01-01 datetime="2000-01-01 00:00:01" class="form-control input-md" value="{{.DataInizio}}"/>
                        </div> 
                    </div>
                    <!-- DATA FINE CONTRATTO --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Data Fine Contratto</label>                          
                        <div class="col-md-4">
                        <input id="datafinecontratto" name="datafinecontratto" type=date min=2000-01-01 datetime="2000-01-01 00:00:01" class="form-control input-md" value="{{.DataFine}}"/>
                        </div> 
                    </div>                                        
                    <!-- RIPARAMETRIZZAZIONE --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Riparametrizzazione</label>                          
                        <div class="col-md-4">
                        <input id="riparametrizzazione" name="riparametrizzazione" type="text" value="{{.Riparametrizzazione}}" class="form-control input-md"/>
                        </div> 
                    </div> 
                    <!-- NUMERO RATE --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Numero Rate</label>                          
                        <div class="col-md-4">
                        <input id="nrate" name="nrate" type="text" value="{{.NRate}}" class="form-control input-md"/>
                        </div> 
                    </div> 
                    <!-- CANONE BASE --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Canone Base</label>                          
                        <div class="col-md-4">
                        <input id="canonebase" name="canonebase" type="text" value="{{.CanoneBase}}" class="form-control input-md"/>
                        </div> 
                    </div> 
                    <!-- CANONE SERVIZI --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Canone Servizi</label>                          
                        <div class="col-md-4">
                        <input id="canoneservizi" name="canoneservizi" type="text" value="{{.CanoneServizi}}" class="form-control input-md"/>
                        </div> 
                    </div> 
                    <!-- CANONE ALTRO --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Canone Altro</label>                          
                        <div class="col-md-4">
                        <input id="canonealtro" name="canonealtro" type="text" class="form-control input-md" value="{{.CanoneAltro}}"/>
                        </div> 
                    </div> 
                    <!-- CANONE TOTALE --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Canone Totale</label>                          
                        <div class="col-md-4">
                        <input id="canonetotale" name="canonetotale" type="text" value="{{.CanoneTot}}" class="form-control input-md"/>
                        </div> 
                    </div>
                    <!-- KM CONTRATTUALI --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">KM Contrattuali</label>                          
                        <div class="col-md-4">
                        <input id="kmcontrattuali" name="kmcontrattuali" type="text" value="{{.KmContrattuali}}" class="form-control input-md"/>
                        </div> 
                    </div>
                    <!-- ADDEBITO KM EXTRA --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Addebito KM Extra</label>                          
                        <div class="col-md-4">
                        <input id="addebitokmextra" name="addebitokmextra" type="text" value="{{.AddebitoKmExtra}}" class="form-control input-md"/>
                        </div> 
                    </div> 
                    <!-- IMPORTO KM --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Importo KM</label>                          
                        <div class="col-md-4">
                        <input id="importokm" name="importokm" type="text" value="{{.ImportoKm}}" class="form-control input-md"/>
                        </div> 
                    </div>
                    <!-- IMPORTO TOTALE--> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Importo Totale</label>                          
                        <div class="col-md-4">
                        <input id="imptotale" name="imptotale" type="text" value="{{.ImportoTot}}" class="form-control input-md"/>
                        </div> 
                    </div>      
                    <!-- KM INIZIO GESTIONE--> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">KM Inizio Gestione</label>                          
                        <div class="col-md-4">
                        <input id="kmingest" name="kmingest" type="text" value="{{.KmInizioGest}}" class="form-control input-md"/>
                        </div> 
                    </div>  
                    <!-- KM FINE GESTIONE--> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">KM Fine Gestione</label>                          
                        <div class="col-md-4">
                        <input id="kmfinegest" name="kmfinegest" type="text" value="{{.KmFineGest}}" class="form-control input-md"/>
                        </div> 
                    </div>                                                                            
                    <!-- NOTE --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Note</label>                          
                        <div class="col-md-4">
                        <input id="note" name="note" type="text" value="{{.Note}}" class="form-control input-md"/>
                        </div> 
                    </div> 
                    <!-- FORNITORI --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Fornitori</label>                          
                        <div class="col-md-4">
                        <input id="fornitore" name="fornitore" type="text" value="{{.Fornitori}}" disabled class="form-control input-md"/>
                        </div> 
                    </div>        
                    <!-- PI FORNITORI --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Partita Iva Fornitore</label>                          
                        <div class="col-md-4">
                        <input id="pifornitore" name="pifornitore" type="text" value="{{.PIFornitori}}" disabled class="form-control input-md"/>
                        </div> 
                    </div>                                                                                                                                       
                    
                    <!-- footer --> 
                    <div class="col-md-12 form-group">
                        <a class="btn btn-primary pull-left" href="http://{{.domainname}}/automezzi/view/fornitori/id!0!id__gte,0">Indietro</a>
                        <button class="btn btn-primary pull-right" type="submit" class="pull-right" value="Continua" />Inserisci</button>
                        
                                  
                    </div>                              

                </form>
        </div>
    </div>
</div>