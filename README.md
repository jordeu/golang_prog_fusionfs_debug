```bash
# Set some variables
bucket_name=umccr-temp-dev
bucket_key_base=stephen/vcfanno_fusionfs_debug

vcfanno_refdata=s3://${bucket_name}/${bucket_key_base}/input/vcfanno_refdata/

nxf_workdir=s3://${bucket_name}/${bucket_key_base}/scratch/
nxf_outdir=s3://${bucket_name}/${bucket_key_base}/output/


# Clone GH repo, upload data
git clone https://github.com/scwatts/vcfanno_fusionfs_debug/ && cd vcfanno_fusionfs_debug/

aws s3 cp data/sample_data/seqc.snvs.vcf.gz s3://${bucket_name}/${bucket_key_base}/input/sample_data/
aws s3 sync data/vcfanno_refdata/ ${vcfanno_refdata}


# Create samplesheet
cat <<EOF > samplesheet.csv
id,vcf
seqc,s3://${bucket_name}/${bucket_key_base}/input/sample_data/seqc.snvs.vcf.gz
EOF


# Run
NXF_VER=23.09.1-edge nextflow -config nextflow_aws.config run vcfanno.nf \
  -ansi-log false \
  -work-dir ${nxf_workdir} \
  --input samplesheet.csv \
  --annotations_dir ${vcfanno_refdata} \
  --outdir ${nxf_outdir}
```
