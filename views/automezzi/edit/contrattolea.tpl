

<div class="col-md-10 content">
    <div class="panel panel-default">
        <div class="panel-heading">
            <h3>Modifica Contratto Leasing</h3>
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
                    <!-- PRIMA RATA --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Prima Rata</label>                          
                        <div class="col-md-4">
                        <input id="prata" name="prata" type="text" value="{{.PRata}}" class="form-control input-md"/>
                        </div> 
                    </div> 
                    <!-- RATE SUCCESSIVE --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Rate Successive</label>                          
                        <div class="col-md-4">
                        <input id="ratesucc" name="ratesucc" type="text" value="{{.RateSucc}}" class="form-control input-md"/>
                        </div> 
                    </div> 
                    <!-- NUMERO RATE --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Numero Rate</label>                          
                        <div class="col-md-4">
                        <input id="nrate" name="nrate" type="text" value="{{.NRate}}" class="form-control input-md"/>
                        </div> 
                    </div> 
                    <!-- RISCATTO --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Riscatto</label>                          
                        <div class="col-md-4">
                        <input id="riscatto" name="riscatto" type="text" value="{{.Riscatto}}" class="form-control input-md"/>
                        </div> 
                    </div> 
                    <!-- DATA RISCATTO --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Data Riscatto</label>                          
                        <div class="col-md-4">
                        <input id="datariscatto" name="datariscatto" type=date min=2000-01-01 datetime="2000-01-01 00:00:01" class="form-control input-md" value="{{.DataRiscatto}}"/>
                        </div> 
                    </div> 
                    <!-- IMPORTO TOTALE --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Importo Totale</label>                          
                        <div class="col-md-4">
                        <input id="importotot" name="importotot" type="text" value="{{.ImportoTot}}" class="form-control input-md"/>
                        </div> 
                    </div>
                    <!-- DATA FINE CONTRATTO --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Data Fine Contratto</label>                          
                        <div class="col-md-4">
                        <input id="datafinecontr" name="datafinecontr" type=date min=2000-01-01 datetime="2000-01-01 00:00:01" class="form-control input-md" value="{{.DataFineContr}}"/>
                        </div> 
                    </div> 
                    <!-- DATA FINE GARANZIA --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Data Fine Garanzia</label>                          
                        <div class="col-md-4">
                        <input id="datafinegaranzia" name="datafinegaranzia" type=date min=2000-01-01 datetime="2000-01-01 00:00:01" class="form-control input-md" value="{{.DataFineGaranzia}}"/>
                        </div> 
                    </div> 
                    <!-- KM INIZIO GESTIONE --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">KM inizio Gestione</label>                          
                        <div class="col-md-4">
                        <input id="kmingest" name="kmingest" type="text" value="{{.KmInizioGest}}" class="form-control input-md"/>
                        </div> 
                    </div>
                    <!-- KM FINE GESTIONE --> 
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
                        <button class="btn btn-primary pull-right" type="submit" value="Continua" />Inserisci</button>
                        
                                  
                    </div>                              

                </form>
        </div>
    </div>
</div>