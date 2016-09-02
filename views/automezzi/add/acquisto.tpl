

<div class="col-md-10 content">
    <div class="panel panel-default">
        <div class="panel-heading">
            <h3>Inserimento Contratto</h3>
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
                    <!-- DATA ACQUISTO --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Data Acquisto</label>                          
                        <div class="col-md-4">
                        <input id="datacont" name="datacont" type=date min=2000-01-01 datetime="2000-01-01 00:00:01" class="form-control input-md" value="{{.DataCont}}" />
                        </div> 
                    </div> 
                    <!-- IMPORTO --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Importo</label>                          
                        <div class="col-md-4">
                        <input id="importo" name="importo" type="text" value="{{.Importo}}" class="form-control input-md"/>
                        </div> 
                    </div> 
                    <!-- AMMORTAMENTO ANNUO --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Ammortamento Annuo</label>                          
                        <div class="col-md-4">
                        <input id="ammortannuo" name="ammortannuo" type="text" value="{{.AmmortamentoAnnuo}}" class="form-control input-md"/>
                        </div> 
                    </div> 
                    <!-- FINE GARANZIA --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Fine Garanzia</label>                          
                        <div class="col-md-4">
                        <input id="finegaranzia" name="finegaranzia" type=date min=2000-01-01 datetime="2000-01-01 00:00:01" class="form-control input-md" value="{{.FineGaranzia}}"/>
                        </div> 
                    </div> 
                    <!-- KM ACQUISTO --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">KM Acquisti</label>                          
                        <div class="col-md-4">
                        <input id="kmacq" name="kmacq" type="text" value="{{.KmAcquisto}}" class="form-control input-md"/>
                        </div> 
                    </div> 
                    <!-- KM INIZIO GESTIONE --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">KM inizio Gestione</label>                          
                        <div class="col-md-4">
                        <input id="kmingest" name="kmingest" type="text" value="{{.KmInizioGest}}" class="form-control input-md"/>
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