

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
                <form class="form-ho rizontal" method="POST">
                    {{ .xsrfdata }}         
                    <!-- NUMERO CONTRATTO --> 
                    <div class="form-group col-lg-12">    
                        <div class="col-md-4">            
                            <label class="control-label" for="nome">Numero Contratto</label>                                    
                            <input id="numcontratto" name="numcontratto" type="text" value="{{.NContratto}}" class="form-control input-md"/>
                        </div> 
                    </div> 
                    <!-- TIPO CONTRATTO -->               
                    <div class="form-group col-lg-12">
                        <label class="control-label" for="contratto">Tipo Contratto</label>  
                        {{if .Errors.TipoContratto}}  {{.Errors.TipoContratto}}  {{end}}
                        <select class="control-label" id="tipocontratto" name="tipocontratto" size="1">
                                <option value="1"selected="selected">Acquisto</option>
                                <option value="2">Leasing</option>
                                <option value="3">Noleggio</option>
                        </select>
                       
                    </div>                                                                                                                                    
                    
                    <!-- footer --> 
                    <div class="col-md-12 form-group">
                        <button class="btn btn-primary pull-right" type="submit" class="pull-right" value="Continua" />Inserisci</button>
                        
                                  
                    </div>                              

                </form>
        </div>
    </div>
</div>