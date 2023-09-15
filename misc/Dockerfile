FROM continuumio/miniconda3:22.11.1

RUN \
  conda install conda-libmamba-solver

RUN \
  echo > ~/.condarc '\
solver: libmamba\n\
channels:\n\
  - umccr\n\
  - bioconda\n\
  - conda-forge\n\
  - defaults'

RUN \
  conda install -y 'vcfanno ==0.3.5' 'bcftools ==1.17' && \
  conda clean -yaf
