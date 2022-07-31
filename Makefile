gen.types:
	find ./schemata | grep .schema.json | while read filepath; do \
	    filename=`echo $${filepath} | sed s/\.schema\.json//g`; \
	    echo $${filename}; \
	    echo $${filepath}; \
		schematyper -o entity/`basename $${filename}`.go --package entity $${filepath}; \
	done ;
